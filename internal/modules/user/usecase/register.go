package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"

	"example-ecommerce/internal/modules/user/domain"

	"github.com/golangid/candi/tracer"
	pbkdf2 "github.com/willy182/goshare/pbkdf2"
)

func (uc *userUsecaseImpl) CreateUser(ctx context.Context, req *domain.RequestUser) (result domain.ResponseUser, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "UserUsecase:CreateUser")
	defer trace.Finish()

	passwordHasher := pbkdf2.NewPBKDF2Hasher(pbkdf2.SaltSize, pbkdf2.SaltSize, pbkdf2.IterationsCount, sha1.New)
	req.Salt = passwordHasher.GenerateSalt()
	err = passwordHasher.ParseSalt(req.Salt)
	if err != nil {
		return
	}

	req.Pin = base64.StdEncoding.EncodeToString(passwordHasher.Hash([]byte(req.Pin)))

	data := req.Deserialize()
	id, err := uc.repoSQL.UserRepo().Save(ctx, &data)
	if err != nil {
		return
	}

	data.ID = id
	result.Serialize(&data)

	return
}
