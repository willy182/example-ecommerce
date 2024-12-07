package usecase

import (
	"context"
	"time"

	"example-ecommerce/internal/modules/order/domain"
	domainUser "example-ecommerce/internal/modules/user/domain"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
	"github.com/google/uuid"
)

func (uc *orderUsecaseImpl) Topup(ctx context.Context, req *domain.RequestTopup) (res domain.ResponseTopup, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:Topup")
	defer trace.Finish()

	user, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.UserID.String()})
	if err != nil {
		return
	}

	amount := float64(req.Amount)

	data := &shareddomain.Order{
		Type:            "TOPUP",
		TransactionType: "DEBIT",
		Status:          "SUCCESS",
		Amount:          amount,
		BalanceBefore:   user.Saldo,
		BalanceAfter:    user.Saldo + amount,
		UserID:          req.UserID,
	}

	var id uuid.UUID
	err = uc.repoSQL.WithTransaction(ctx, func(ctxTx context.Context) error {
		resID, e := uc.repoSQL.OrderRepo().Save(ctxTx, data)
		if e != nil {
			return e
		}

		user.Saldo += amount
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
	res.BalanceBefore = user.Saldo - amount
	res.BalanceAfter = user.Saldo
	res.CreatedAt = time.Now().Format(time.DateTime)
	return
}
