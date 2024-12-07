package domain

import (
	"time"

	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/google/uuid"
)

// ResponseUserList model
type ResponseUserList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseUser   `json:"data"`
}

// ResponseUser model
type ResponseUser struct {
	ID          uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Saldo       float64   `json:"saldo,omitempty"`
	Address     string    `json:"address"`
	CreatedAt   string    `json:"created_date"`
}

// ResponseLogin model
type ResponseLogin struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// Serialize from db model
func (r *ResponseUser) Serialize(source *shareddomain.User) {
	r.ID = source.ID
	r.FirstName = source.FirstName
	r.LastName = source.LastName
	r.PhoneNumber = source.PhoneNumber
	r.Saldo = source.Saldo
	r.Address = source.Address
	r.CreatedAt = source.CreatedAt.Format(time.DateTime)
}
