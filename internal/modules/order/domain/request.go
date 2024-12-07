package domain

import (
	"github.com/google/uuid"
)

// RequestTopup model
type RequestTopup struct {
	UserID uuid.UUID
	Amount int `json:"amount"`
}

// RequestPayment model
type RequestPayment struct {
	UserID  uuid.UUID
	Amount  int    `json:"amount"`
	Remarks string `json:"remarks"`
}

// RequestTransfer model
type RequestTransfer struct {
	UserID     uuid.UUID
	Amount     int       `json:"amount"`
	Remarks    string    `json:"remarks"`
	TargetUser uuid.UUID `json:"target_user"`
}

// ReceiveTransfer model
type ReceiveTransfer struct {
	UserID  uuid.UUID
	Amount  int    `json:"amount"`
	Remarks string `json:"remarks"`
}
