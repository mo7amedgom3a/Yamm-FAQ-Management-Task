package dto

import "github.com/google/uuid"

type CreateStoreRequest struct {
	Name       string    `json:"name"`
	MerchantID uuid.UUID `json:"merchant_id"`
}

type StoreResponse struct {
	ID         uuid.UUID     `json:"id"`
	Name       string        `json:"name"`
	MerchantID uuid.UUID     `json:"merchant_id"`
	FAQs       []FAQResponse `json:"faqs"`
}
