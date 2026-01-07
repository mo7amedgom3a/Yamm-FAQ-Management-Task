package models

import (
	"time"

	"github.com/google/uuid"
)

type FAQ struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	CategoryID uuid.UUID `gorm:"type:uuid;not null"`
	Category   FAQCategory

	StoreID *uuid.UUID `gorm:"type:uuid"`
	Store   *Store

	IsGlobal bool `gorm:"not null;default:false"`

	CreatedBy uuid.UUID `gorm:"type:uuid;not null"`
	Creator   User      `gorm:"foreignKey:CreatedBy"`

	Translations []FAQTranslation `gorm:"foreignKey:FAQID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
