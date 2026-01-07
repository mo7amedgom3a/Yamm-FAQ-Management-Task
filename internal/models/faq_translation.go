package models

import (
	"time"

	"github.com/google/uuid"
)

type FAQTranslation struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	FAQID        uuid.UUID `gorm:"type:uuid;not null"`
	LanguageCode string    `gorm:"type:varchar(5);not null"`
	Question     string    `gorm:"type:text;not null"`
	Answer       string    `gorm:"type:text;not null"`

	FAQ FAQ `gorm:"constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
