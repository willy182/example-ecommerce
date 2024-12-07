package domain

import (
	"time"

	"github.com/google/uuid"
)

// Order model
type Order struct {
	ID                 uuid.UUID  `gorm:"column:id;type:uuid;primary_key" json:"id"`
	Type               string     `gorm:"column:type;type:varchar(20)" json:"type"`
	TransactionType    string     `gorm:"column:transaction_type;type:varchar(20)" json:"transaction_type"`
	Status             string     `gorm:"column:status;type:varchar(20)" json:"status"`
	Amount             float64    `gorm:"column:amount;type:decimal(10,2)" json:"amount"`
	BalanceBefore      float64    `gorm:"column:balance_before;type:decimal(10,2)" json:"balance_before"`
	BalanceAfter       float64    `gorm:"column:balance_after;type:decimal(10,2)" json:"balance_after"`
	Remarks            string     `gorm:"column:remarks;type:varchar(100)" json:"remarks"`
	UserID             uuid.UUID  `gorm:"column:user_id;type:uuid" json:"user_id"`
	TransferToUserID   *uuid.UUID `gorm:"column:transfer_to_user_id;type:uuid" json:"transfer_to_user_id"`
	CreatedAt          time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time  `gorm:"column:updated_at" json:"updated_at"`
	UserData           User       `gorm:"foreignKey:UserID"`
	TransferToUserData User       `gorm:"foreignKey:TransferToUserID"`
}

// TableName return table name of Order model
func (Order) TableName() string {
	return "orders"
}
