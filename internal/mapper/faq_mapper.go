package mapper

import (
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FAQMapper struct {
	translationMapper *FAQTranslationMapper
}

func NewFAQMapper(translationMapper *FAQTranslationMapper) *FAQMapper {
	return &FAQMapper{
		translationMapper: translationMapper,
	}
}

func (m *FAQMapper) ToModel(req dto.CreateFAQRequest) models.FAQ {
	return models.FAQ{
		CategoryID: req.CategoryID,
		IsGlobal:   req.IsGlobal,
		StoreID:    req.StoreID,
	}
}

func (m *FAQMapper) ToDTO(faq models.FAQ) dto.FAQResponse {
	return dto.FAQResponse{
		ID:           faq.ID,
		CategoryID:   faq.CategoryID,
		IsGlobal:     faq.IsGlobal,
		StoreID:      faq.StoreID,
		Translations: m.translationMapper.ToDTOs(faq.Translations),
	}
}

func (m *FAQMapper) ToDTOs(faqs []models.FAQ) []dto.FAQResponse {
	dtos := make([]dto.FAQResponse, len(faqs))
	for i, faq := range faqs {
		dtos[i] = m.ToDTO(faq)
	}
	return dtos
}
