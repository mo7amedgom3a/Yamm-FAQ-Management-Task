package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleMerchant UserRole = "merchant"
	RoleCustomer UserRole = "customer"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         UserRole  `gorm:"type:varchar(20);not null"`

	Store *Store `gorm:"foreignKey:MerchantID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
