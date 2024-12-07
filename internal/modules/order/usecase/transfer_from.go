package usecase

import (
	"context"
	"fmt"

	"example-ecommerce/internal/modules/order/domain"
	domainUser "example-ecommerce/internal/modules/user/domain"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
	"github.com/google/uuid"
)

func (uc *orderUsecaseImpl) TransferFrom(ctx context.Context, req *domain.ReceiveTransfer) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:TransferFrom")
	defer trace.Finish()

	user, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.UserID.String()})
	if err != nil {
		return
	}

	amount := float64(req.Amount)

	if user.Saldo-amount < 0 {
		err = fmt.Errorf("Saldo anda tidak cukup untuk melakukan transfer.")
		return
	}

	userTarget, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.Remarks})
	if err != nil {
		return
	}

	userID, _ := uuid.Parse(req.Remarks)
	data := &shareddomain.Order{
		Type:            "RECEIVE_TRANSFER",
		TransactionType: "DEBIT",
		Status:          "SUCCESS",
		Amount:          amount,
		BalanceBefore:   userTarget.Saldo,
		BalanceAfter:    userTarget.Saldo + amount,
		UserID:          userID,
		Remarks:         fmt.Sprintf("Transfer from %s %s", user.FirstName, user.LastName),
	}

	err = uc.repoSQL.WithTransaction(ctx, func(ctxTx context.Context) error {
		_, e := uc.repoSQL.OrderRepo().Save(ctxTx, data)
		if e != nil {
			return e
		}

		user.Saldo -= amount
		_, e = uc.repoSQL.UserRepo().Save(ctxTx, &user, candishared.DBUpdateSetUpdatedFields("Saldo"))
		if e != nil {
			return e
		}

		userTarget.Saldo += amount
		_, e = uc.repoSQL.UserRepo().Save(ctxTx, &userTarget, candishared.DBUpdateSetUpdatedFields("Saldo"))
		if e != nil {
			return e
		}

		return nil
	})
	if err != nil {
		return
	}

	return
}
