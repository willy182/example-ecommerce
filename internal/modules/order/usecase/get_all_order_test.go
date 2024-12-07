package usecase

import (
	"context"
	"errors"
	"testing"

	"example-ecommerce/internal/modules/order/domain"
	mockrepo "example-ecommerce/pkg/mocks/modules/order/repository"
	mocksharedrepo "example-ecommerce/pkg/mocks/shared/repository"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_GetAllOrder(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Order{}, nil)
		orderRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)

		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllOrder(context.Background(), &domain.FilterOrder{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("FetchAll", mock.Anything, mock.Anything, mock.Anything).Return([]shareddomain.Order{}, errors.New("Error"))
		orderRepo.On("Count", mock.Anything, mock.Anything).Return(10)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)

		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetAllOrder(context.Background(), &domain.FilterOrder{})
		assert.Error(t, err)
	})
}
