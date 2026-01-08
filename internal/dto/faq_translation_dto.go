package dto

import "github.com/google/uuid"

type CreateFAQTranslationRequest struct {
	FAQID    uuid.UUID
	Language string
	Question string
	Answer   string
}

type FAQTranslationResponse struct {
	ID       uuid.UUID
	FAQID    uuid.UUID
	Language string
	Question string
	Answer   string
}
