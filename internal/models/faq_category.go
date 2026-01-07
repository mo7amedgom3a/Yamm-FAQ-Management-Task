package models

import (
	"time"

	"github.com/google/uuid"
)

type FAQCategory struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"not null"`

	FAQs []FAQ `gorm:"foreignKey:CategoryID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
