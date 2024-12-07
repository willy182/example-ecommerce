package usecase

import (
	"context"

	"example-ecommerce/internal/modules/order/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *orderUsecaseImpl) GetAllOrder(ctx context.Context, filter *domain.FilterOrder) (result domain.ResponseOrderList, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "OrderUsecase:GetAllOrder")
	defer trace.Finish()

	data, err := uc.repoSQL.OrderRepo().FetchAll(ctx, filter)
	if err != nil {
		return result, err
	}
	count := uc.repoSQL.OrderRepo().Count(ctx, filter)
	result.Meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	result.Data = make([]domain.ResponseOrder, len(data))
	for i, detail := range data {
		result.Data[i].Serialize(&detail)
	}

	return
}
