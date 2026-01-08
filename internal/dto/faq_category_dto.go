package dto

type CreateFAQCategoryRequest struct {
	Name string `json:"name"`
}

type FAQCategoryResponse struct {
	ID   string        `json:"id"`
	Name string        `json:"name"`
	FAQs []FAQResponse `json:"faqs"`
}
