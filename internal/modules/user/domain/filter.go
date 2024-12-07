package domain

import "github.com/golangid/candi/candishared"

// FilterUser model
type FilterUser struct {
	candishared.Filter
	ID          string   `json:"id"`
	PhoneNumber string   `json:"phone_number"`
	Preloads    []string `json:"-"`
}
