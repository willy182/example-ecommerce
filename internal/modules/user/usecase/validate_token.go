package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"example-ecommerce/configs/rsa"
	"example-ecommerce/internal/modules/user/domain"
	"example-ecommerce/pkg/helper"

	oldJWT "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

// ValidateToken implement TokenValidator
func (uc *userUsecaseImpl) ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:ValidateToken")

	var err error
	defer func() {
		helper.RecoverPanic(&err)
		trace.SetTag("token", token)
		trace.Finish(tracer.FinishWithError(err))
	}()

	_, _, err = uc.verifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	resultTokenClaim, _, err := uc.validateToken(ctx, token)
	if err != nil {
		return nil, err
	}

	tc := resultTokenClaim.(domain.TokenClaims)

	standardClaim := oldJWT.StandardClaims{
		ExpiresAt: int64(tc.ExpiresAt),
		Id:        tc.Jti,
		IssuedAt:  int64(tc.IssuedAt),
		Subject:   tc.Subject,
	}

	additional := domain.AdditionalClaims{
		PhoneNumber: tc.PhoneNumber,
	}

	return &candishared.TokenClaim{
		StandardClaims: standardClaim,
		Additional:     additional,
	}, nil
}

func (uc *userUsecaseImpl) verifyToken(ctx context.Context, token string) (interface{}, int, error) {
	trace, _ := tracer.StartTraceWithContext(ctx, "UserUsecase:verifyToken")

	var err error
	defer func() {
		helper.RecoverPanic(&err)
		trace.SetTag("token", token)
		trace.Finish(tracer.FinishWithError(err))
	}()

	tokenClaim := jwt.MapClaims{}

	jwtResult, err := jwt.ParseWithClaims(token, tokenClaim, func(tkn *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != tkn.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if (jwtResult == nil && err != nil) || len(tokenClaim) == 0 {
		return nil, http.StatusUnauthorized, err
	}

	expired, ok := tokenClaim["exp"].(float64)
	if !ok {
		err = fmt.Errorf("invalid token")
		return nil, http.StatusUnauthorized, err
	}

	nowDate, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	expiredDate, err := time.Parse(time.RFC3339, time.Unix(int64(expired), 0).Format(time.RFC3339))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// check token expired
	if nowDate.After(expiredDate) {
		err = fmt.Errorf("token has expired")
		return nil, http.StatusUnauthorized, err
	}

	return tokenClaim, http.StatusOK, nil
}

func (uc *userUsecaseImpl) validateToken(ctx context.Context, oldToken string) (interface{}, int, error) {
	trace, _ := tracer.StartTraceWithContext(ctx, "UserUsecase:validateToken")

	var err error
	defer func() {
		helper.RecoverPanic(&err)
		trace.SetTag("oldToken", oldToken)
		trace.Finish(tracer.FinishWithError(err))
	}()

	pubKey, err := rsa.InitPublicKey()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(oldToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if token == nil {
		err = fmt.Errorf("invalid token")
		return nil, http.StatusUnauthorized, err
	}

	var tokenClaim domain.TokenClaims
	tokenClaim.ExpiresAt = claims["exp"].(float64)
	tokenClaim.Jti = claims["jti"].(string)
	tokenClaim.IssuedAt = claims["iat"].(float64)
	tokenClaim.Subject = claims["sub"].(string)
	tokenClaim.PhoneNumber = claims["phone_number"].(string)

	return tokenClaim, http.StatusOK, nil
}
