package models

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"not null"`
	MerchantID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`

	Merchant User `gorm:"constraint:OnDelete:CASCADE"`

	FAQs []FAQ `gorm:"foreignKey:StoreID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
