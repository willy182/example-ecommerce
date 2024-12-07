package usecase

import (
	"context"
	"errors"
	"testing"

	"example-ecommerce/internal/modules/order/domain"
	mockrepo "example-ecommerce/pkg/mocks/modules/order/repository"
	mockrepouser "example-ecommerce/pkg/mocks/modules/user/repository"
	mocksharedrepo "example-ecommerce/pkg/mocks/shared/repository"
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_orderUsecaseImpl_Transfer(t *testing.T) {
	ctx := context.Background()
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		userRepo := &mockrepouser.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)

		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Save", mock.Anything, mock.Anything).Return(uuid.New(), nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)
		repoSQL.On("UserRepo").Return(userRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.Transfer(ctx, &domain.RequestTransfer{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		userRepo := &mockrepouser.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, errors.New("Error"))
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)

		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Save", mock.Anything, mock.Anything).Return(uuid.New(), nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)
		repoSQL.On("UserRepo").Return(userRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.Transfer(ctx, &domain.RequestTransfer{})
		assert.Error(t, err)
	})

	t.Run("Testcase #3: Negative", func(t *testing.T) {
		userRepo := &mockrepouser.UserRepository{}
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)
		userRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.User{}, nil)

		orderRepo := &mockrepo.OrderRepository{}
		orderRepo.On("Save", mock.Anything, mock.Anything).Return(uuid.New(), errors.New("Error"))

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("OrderRepo").Return(orderRepo)
		repoSQL.On("UserRepo").Return(userRepo)
		repoSQL.On("WithTransaction", mock.Anything,
			mock.AnythingOfType("func(context.Context) error")).
			Return(errors.New("Error")).
			Run(func(args mock.Arguments) {
				arg := args.Get(1).(func(context.Context) error)
				arg(ctx)
			})
		uc := orderUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.Transfer(ctx, &domain.RequestTransfer{})
		assert.Error(t, err)
	})
}