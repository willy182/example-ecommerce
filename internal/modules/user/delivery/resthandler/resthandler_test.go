// Code generated by candi v1.18.4.

package resthandler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example-ecommerce/internal/modules/user/domain"
	mockusecase "example-ecommerce/pkg/mocks/modules/user/usecase"
	mocksharedusecase "example-ecommerce/pkg/mocks/shared/usecase"

	"github.com/golangid/candi/candihelper"
	mockdeps "github.com/golangid/candi/mocks/codebase/factory/dependency"
	mockinterfaces "github.com/golangid/candi/mocks/codebase/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCase struct {
	name, reqBody                       string
	wantValidateError, wantUsecaseError error
	wantRespCode                        int
}

var (
	errFoo = errors.New("something error")
)

func TestNewRestHandler(t *testing.T) {
	mockMiddleware := &mockinterfaces.Middleware{}
	mockMiddleware.On("HTTPPermissionACL", mock.Anything).Return(func(http.Handler) http.Handler { return nil })
	mockValidator := &mockinterfaces.Validator{}

	mockDeps := &mockdeps.Dependency{}
	mockDeps.On("GetMiddleware").Return(mockMiddleware)
	mockDeps.On("GetValidator").Return(mockValidator)

	handler := NewRestHandler(nil, mockDeps)
	assert.NotNil(t, handler)

	mockRoute := &mockinterfaces.RESTRouter{}
	mockRoute.On("Group", mock.Anything, mock.Anything).Return(mockRoute)
	mockRoute.On("GET", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("POST", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("PUT", mock.Anything, mock.Anything, mock.Anything)
	mockRoute.On("DELETE", mock.Anything, mock.Anything, mock.Anything)
	handler.Mount(mockRoute)
}

func TestRestHandler_createUser(t *testing.T) {
	tests := []testCase{
		{
			name: "Testcase #1: Positive", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: nil, wantRespCode: http.StatusCreated,
		},
		{
			name: "Testcase #2: Negative", reqBody: `{"email": test@test.com}`, wantUsecaseError: nil, wantRespCode: http.StatusBadRequest,
		},
		{
			name: "Testcase #3: Negative", reqBody: `{"email": "test@test.com"}`, wantUsecaseError: errFoo, wantRespCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userUsecase := &mockusecase.UserUsecase{}
			userUsecase.On("CreateUser", mock.Anything, mock.Anything).Return(domain.ResponseUser{}, tt.wantUsecaseError)
			mockValidator := &mockinterfaces.Validator{}
			mockValidator.On("ValidateDocument", mock.Anything, mock.Anything).Return(tt.wantValidateError)

			uc := &mocksharedusecase.Usecase{}
			uc.On("User").Return(userUsecase)

			handler := RestHandler{uc: uc, validator: mockValidator}

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.reqBody))
			req.Header.Add(candihelper.HeaderContentType, candihelper.HeaderMIMEApplicationJSON)
			res := httptest.NewRecorder()
			handler.createUser(res, req)
			assert.Equal(t, tt.wantRespCode, res.Code)
		})
	}
}