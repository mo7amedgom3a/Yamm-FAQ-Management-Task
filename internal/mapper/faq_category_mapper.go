package mapper

import (
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FAQCategoryMapper struct {
	faqMapper *FAQMapper
}

func NewFAQCategoryMapper(faqMapper *FAQMapper) *FAQCategoryMapper {
	return &FAQCategoryMapper{
		faqMapper: faqMapper,
	}
}

func (m *FAQCategoryMapper) ToModel(req dto.CreateFAQCategoryRequest) models.FAQCategory {
	return models.FAQCategory{
		Name: req.Name,
	}
}

func (m *FAQCategoryMapper) ToDTO(category models.FAQCategory) dto.FAQCategoryResponse {
	faqs := make([]dto.FAQResponse, 0)
	if len(category.FAQs) > 0 && m.faqMapper != nil {
		faqs = m.faqMapper.ToDTOs(category.FAQs)
	}
	return dto.FAQCategoryResponse{
		ID:   category.ID.String(),
		Name: category.Name,
		FAQs: faqs,
	}
}

func (m *FAQCategoryMapper) ToDTOs(categories []models.FAQCategory) []dto.FAQCategoryResponse {
	dtos := make([]dto.FAQCategoryResponse, len(categories))
	for i, category := range categories {
		dtos[i] = m.ToDTO(category)
	}
	return dtos
}
