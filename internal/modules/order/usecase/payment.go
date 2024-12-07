package usecase

import (
	"context"
	"fmt"
	"time"

	"example-ecommerce/internal/modules/order/domain"
	domainUser "example-ecommerce/internal/modules/user/domain"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
	"github.com/google/uuid"
)

func (uc *orderUsecaseImpl) Payment(ctx context.Context, req *domain.RequestPayment) (res domain.ResponsePayment, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:Payment")
	defer trace.Finish()

	user, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.UserID.String()})
	if err != nil {
		return
	}

	amount := float64(req.Amount)

	if user.Saldo-amount < 0 {
		err = fmt.Errorf("Saldo anda tidak cukup untuk melakukan pembayaran.")
		return
	}

	data := &shareddomain.Order{
		Type:            "PAYMENT",
		TransactionType: "CREDIT",
		Status:          "SUCCESS",
		Amount:          amount,
		BalanceBefore:   user.Saldo,
		BalanceAfter:    user.Saldo - amount,
		UserID:          req.UserID,
		Remarks:         req.Remarks,
	}

	var id uuid.UUID
	err = uc.repoSQL.WithTransaction(ctx, func(ctxTx context.Context) error {
		resID, e := uc.repoSQL.OrderRepo().Save(ctxTx, data)
		if e != nil {
			return e
		}

		user.Saldo -= amount
		_, e = uc.repoSQL.UserRepo().Save(ctxTx, &user, candishared.DBUpdateSetUpdatedFields("Saldo"))
		if e != nil {
			return e
		}

		id = resID

		return nil
	})
	if err != nil {
		return
	}

	res.ID = id
	res.Amount = float64(req.Amount)
	res.Remarks = req.Remarks
	res.BalanceBefore = user.Saldo + amount
	res.BalanceAfter = user.Saldo
	res.CreatedAt = time.Now().Format(time.DateTime)
	return
}
