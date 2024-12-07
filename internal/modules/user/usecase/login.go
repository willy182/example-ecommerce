package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"example-ecommerce/configs/rsa"
	"example-ecommerce/internal/modules/user/domain"
	"example-ecommerce/pkg/helper"
	"example-ecommerce/pkg/shared"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golangid/candi/tracer"
	"github.com/google/uuid"
	pbkdf2 "github.com/willy182/goshare/pbkdf2"
	goshared "github.com/willy182/goshare/shared"
)

func (uc *userUsecaseImpl) Login(ctx context.Context, data *domain.LoginUser) (res domain.ResponseLogin, statusCode int, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:Login")
	defer trace.Finish()

	repoFilter := domain.FilterUser{PhoneNumber: data.PhoneNumber}
	existing, err := uc.repoSQL.UserRepo().Find(ctx, &repoFilter)
	if err != nil {
		err = fmt.Errorf("Phone Number and PIN doesn't match.")
		statusCode = http.StatusBadRequest
		return
	}

	passwordHasher := pbkdf2.NewPBKDF2Hasher(pbkdf2.SaltSize, pbkdf2.SaltSize, pbkdf2.IterationsCount, sha1.New)
	passwordHasher.ParseSalt(existing.Salt)
	base64Data := base64.StdEncoding.EncodeToString(passwordHasher.Hash([]byte(data.Pin)))

	if existing.Pin != base64Data {
		statusCode = http.StatusUnauthorized
		err = fmt.Errorf("Phone Number and PIN doesn't match.")
		return
	}

	claims := domain.Claim{}
	claims.Subject = existing.ID.String()
	claims.PhoneNumber = existing.PhoneNumber

	accessToken, err := uc.generateAccessToken(ctx, claims)
	if err != nil {
		statusCode = http.StatusBadRequest
		return
	}

	refreshToken := goshared.RandomStringBase64(46)

	statusCode = http.StatusOK

	res.AccessToken = accessToken.AccessToken
	res.RefreshToken = refreshToken
	return
}

func (uc *userUsecaseImpl) generateAccessToken(ctx context.Context, cl domain.Claim) (res domain.AccessToken, err error) {
	trace, _ := tracer.StartTraceWithContext(ctx, "UserUsecase:generateAccessToken")
	defer func() {
		helper.RecoverPanic(&err)
		trace.Log("accessToken", res)
		trace.Finish(tracer.FinishWithError(err))
	}()

	signKey, err := rsa.InitPrivateKey()
	if err != nil {
		return
	}

	tokenAge, err := time.ParseDuration(shared.GetEnv().AccessTokenAge)
	if err != nil {
		return
	}

	now := time.Now()
	now = helper.ToAsiaJakartaTime(now)
	age := now.Add(tokenAge)

	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)

	mixid := cl.PhoneNumber + "----" + uuid.NewString()
	jti := helper.GenerateTokenByString(mixid)

	claims["exp"] = age.Unix()
	claims["jti"] = jti
	claims["iat"] = now.Unix()
	claims["sub"] = cl.Subject
	claims["phone_number"] = cl.PhoneNumber
	token.Claims = claims

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return
	}

	res.AccessToken = tokenString
	res.ExpiredAt = age

	return
}
