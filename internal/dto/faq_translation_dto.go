package dto

import "github.com/google/uuid"

type CreateFAQTranslationRequest struct {
	FAQID    uuid.UUID `json:"faq_id"`
	Language string    `json:"language"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}

type FAQTranslationResponse struct {
	ID       uuid.UUID `json:"id"`
	FAQID    uuid.UUID `json:"faq_id"`
	Language string    `json:"language"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}
