package mapper

import (
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FAQCategoryMapper struct{}

func NewFAQCategoryMapper() *FAQCategoryMapper {
	return &FAQCategoryMapper{}
}

func (m *FAQCategoryMapper) ToModel(req dto.CreateFAQCategoryRequest) models.FAQCategory {
	return models.FAQCategory{
		Name: req.Name,
	}
}

func (m *FAQCategoryMapper) ToDTO(category models.FAQCategory) dto.FAQCategoryResponse {
	return dto.FAQCategoryResponse{
		ID:   category.ID.String(),
		Name: category.Name,
		// FAQs should be mapped separately if needed, or we can add a helper here
		// For now, leaving FAQs empty or handling it in the service if complex mapping is needed
		// But usually the mapper should handle it.
		// Let's assume simple mapping for now.
	}
}

func (m *FAQCategoryMapper) ToDTOs(categories []models.FAQCategory) []dto.FAQCategoryResponse {
	dtos := make([]dto.FAQCategoryResponse, len(categories))
	for i, category := range categories {
		dtos[i] = m.ToDTO(category)
	}
	return dtos
}
