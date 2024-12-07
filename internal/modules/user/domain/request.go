package domain

import (
	shareddomain "example-ecommerce/pkg/shared/domain"

	"github.com/google/uuid"
)

// RequestUser model
type RequestUser struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Pin         string    `json:"pin"`
	Salt        string    `json:"-"`
	Saldo       float64   `json:"saldo,omitempty"`
	Address     string    `json:"address"`
}

// LoginUser model
type LoginUser struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}

// Deserialize to db model
func (r *RequestUser) Deserialize() (res shareddomain.User) {
	res.FirstName = r.FirstName
	res.LastName = r.LastName
	res.PhoneNumber = r.PhoneNumber
	res.Pin = r.Pin
	res.Salt = r.Salt
	res.Saldo = r.Saldo
	res.Address = r.Address
	return
}
