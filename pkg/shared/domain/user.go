package domain

import (
	"time"

	"github.com/google/uuid"
)

// User model
type User struct {
	ID          uuid.UUID `gorm:"column:id;primary_key" json:"id"`
	FirstName   string    `gorm:"column:first_name;type:varchar(100)" json:"first_name"`
	LastName    string    `gorm:"column:last_name;type:varchar(100)" json:"last_name"`
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(16)" json:"phone_number"`
	Pin         string    `gorm:"column:pin;type:varchar(255)" json:"pin"`
	Salt        string    `gorm:"column:salt;type:varchar(255)" json:"-"`
	Saldo       float64   `gorm:"column:saldo;type:decimal(10,2)" json:"saldo"`
	Address     string    `gorm:"column:address;type:dtext" json:"address"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName return table name of User model
func (User) TableName() string {
	return "users"
}
