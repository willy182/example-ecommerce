package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"example-ecommerce/internal/modules/order/domain"
	domainUser "example-ecommerce/internal/modules/user/domain"
	"example-ecommerce/pkg/shared"
	shareddomain "example-ecommerce/pkg/shared/domain"

	taskqueueworker "github.com/golangid/candi/codebase/app/task_queue_worker"
	"github.com/golangid/candi/tracer"
	"github.com/google/uuid"
)

func (uc *orderUsecaseImpl) Transfer(ctx context.Context, req *domain.RequestTransfer) (res domain.ResponseTransfer, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:Transfer")
	defer trace.Finish()

	user, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.UserID.String()})
	if err != nil {
		return
	}

	userTarget, err := uc.repoSQL.UserRepo().Find(ctx, &domainUser.FilterUser{ID: req.TargetUser.String()})
	if err != nil {
		err = fmt.Errorf("Target user not found.")
		return
	}

	amount := float64(req.Amount)

	if user.Saldo-amount < 0 {
		err = fmt.Errorf("Saldo anda tidak cukup untuk melakukan transfer.")
		return
	}

	data := &shareddomain.Order{
		Type:             "TRANSFER",
		TransactionType:  "CREDIT",
		Status:           "SUCCESS",
		Amount:           amount,
		BalanceBefore:    user.Saldo,
		BalanceAfter:     user.Saldo - amount,
		UserID:           req.UserID,
		Remarks:          req.Remarks,
		TransferToUserID: &userTarget.ID,
	}

	var id uuid.UUID
	err = uc.repoSQL.WithTransaction(ctx, func(ctxTx context.Context) error {
		resID, e := uc.repoSQL.OrderRepo().Save(ctxTx, data)
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
	res.BalanceBefore = user.Saldo
	res.BalanceAfter = user.Saldo - amount
	res.CreatedAt = time.Now().Format(time.DateTime)

	param := domain.ReceiveTransfer{
		UserID:  req.UserID,
		Amount:  req.Amount,
		Remarks: req.TargetUser.String(),
	}

	paramByte, _ := json.Marshal(param)

	_, _ = taskqueueworker.AddJob(ctx, &taskqueueworker.AddJobRequest{
		TaskName: "TRANSFER",
		MaxRetry: shared.GetEnv().RetryCount,
		Args:     paramByte,
	})

	return
}
