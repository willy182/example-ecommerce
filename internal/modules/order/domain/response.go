package domain

import (
	shareddomain "example-ecommerce/pkg/shared/domain"
	"strconv"
	"time"

	"github.com/golangid/candi/candishared"
	"github.com/google/uuid"
)

// ResponseOrderList model
type ResponseOrderList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseOrder  `json:"data"`
}

// ResponseOrder model
type ResponseOrder struct {
	ID              uuid.UUID `json:"id"`
	Status          string    `json:"Status"`
	UserID          uuid.UUID `json:"user_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          string    `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   string    `json:"balance_before"`
	BalanceAfter    string    `json:"balance_after"`
	CreatedAt       string    `json:"created_date"`
}

// Serialize from db model
func (r *ResponseOrder) Serialize(source *shareddomain.Order) {
	r.ID = source.ID
	r.Status = source.Status
	r.UserID = source.UserID
	r.TransactionType = source.TransactionType
	r.Amount = strconv.Itoa(int(source.Amount))
	r.Remarks = source.Remarks
	r.BalanceBefore = strconv.Itoa(int(source.BalanceBefore))
	r.BalanceAfter = strconv.Itoa(int(source.BalanceAfter))
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
}

// ResponseTopup model
type ResponseTopup struct {
	ID            uuid.UUID `json:"top_up_id"`
	Amount        float64   `json:"amount_top_up"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedAt     string    `json:"created_date"`
}

// ResponsePayment model
type ResponsePayment struct {
	ID            uuid.UUID `json:"payment_id"`
	Amount        float64   `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedAt     string    `json:"created_date"`
}

// ResponseTransfer model
type ResponseTransfer struct {
	ID            uuid.UUID `json:"transfer_id"`
	Amount        float64   `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedAt     string    `json:"created_date"`
}
