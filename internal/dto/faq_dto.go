package dto

import "github.com/google/uuid"

type CreateFAQRequest struct {
	CategoryID uuid.UUID  `json:"category_id"`
	IsGlobal   bool       `json:"is_global"`
	StoreID    *uuid.UUID `json:"store_id"`
	CreatedBy  string     `json:"-"` // Set by handler
}

type FAQResponse struct {
	ID           uuid.UUID                `json:"id"`
	CategoryID   uuid.UUID                `json:"category_id"`
	IsGlobal     bool                     `json:"is_global"`
	StoreID      *uuid.UUID               `json:"store_id"`
	Translations []FAQTranslationResponse `json:"translations"`
}
