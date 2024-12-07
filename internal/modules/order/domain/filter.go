package domain

import "github.com/golangid/candi/candishared"

// FilterOrder model
type FilterOrder struct {
	candishared.Filter
	ID              string   `json:"id"`
	UserID          string   `json:"user_id"`
	TransactionType string   `json:"type_transaction"`
	Preloads        []string `json:"-"`
}
