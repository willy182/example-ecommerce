// Code generated by candi v1.18.4.

package repository

import (
	"context"

	"example-ecommerce/internal/modules/order/domain"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/google/uuid"
)

// OrderRepository abstract interface
type OrderRepository interface {
	FetchAll(ctx context.Context, filter *domain.FilterOrder) ([]shareddomain.Order, error)
	Count(ctx context.Context, filter *domain.FilterOrder) int
	Find(ctx context.Context, filter *domain.FilterOrder) (shareddomain.Order, error)
	Save(ctx context.Context, data *shareddomain.Order, updateOptions ...candishared.DBUpdateOptionFunc) (id uuid.UUID, err error)
	Delete(ctx context.Context, filter *domain.FilterOrder) (err error)
}
