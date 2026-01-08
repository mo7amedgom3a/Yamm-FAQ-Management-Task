package dto

import "github.com/google/uuid"

type CreateFAQRequest struct {
	CategoryID uuid.UUID
	IsGlobal   bool
	StoreID    *uuid.UUID
}

type FAQResponse struct {
	ID           uuid.UUID
	CategoryID   uuid.UUID
	IsGlobal     bool
	StoreID      *uuid.UUID
	Translations []FAQTranslationResponse
}
