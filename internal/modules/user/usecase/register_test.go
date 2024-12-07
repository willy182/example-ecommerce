package usecase

import (
	"context"
	"errors"
	"testing"

	"example-ecommerce/internal/modules/user/domain"
	mockrepo "example-ecommerce/pkg/mocks/modules/user/repository"
	mocksharedrepo "example-ecommerce/pkg/mocks/shared/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_userUsecaseImpl_CreateUser(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {
		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Save", mock.Anything, mock.Anything).Return(uuid.New(), nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateUser(context.Background(), &domain.RequestUser{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		userRepo := &mockrepo.UserRepository{}
		userRepo.On("Save", mock.Anything, mock.Anything).Return(uuid.New(), errors.New("error"))

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("UserRepo").Return(userRepo)

		uc := userUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.CreateUser(context.Background(), &domain.RequestUser{})
		assert.Error(t, err)
	})
}
