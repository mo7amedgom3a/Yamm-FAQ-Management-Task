package dto

type CreateFAQCategoryRequest struct {
	Name string
}

type FAQCategoryResponse struct {
	ID   string
	Name string
	FAQs []FAQResponse
}