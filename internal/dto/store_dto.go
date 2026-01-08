package dto

import "github.com/google/uuid"

type CreateStoreRequest struct {
	Name       string
	MerchantID uuid.UUID
}

type StoreResponse struct {
	ID         uuid.UUID
	Name       string
	MerchantID uuid.UUID
	FAQs       []FAQResponse
}
