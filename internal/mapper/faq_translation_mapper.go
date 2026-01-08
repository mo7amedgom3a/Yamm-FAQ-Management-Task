package mapper

import (
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FAQTranslationMapper struct{}

func NewFAQTranslationMapper() *FAQTranslationMapper {
	return &FAQTranslationMapper{}
}

func (m *FAQTranslationMapper) ToModel(req dto.CreateFAQTranslationRequest) models.FAQTranslation {
	return models.FAQTranslation{
		FAQID:        req.FAQID,
		LanguageCode: req.Language,
		Question:     req.Question,
		Answer:       req.Answer,
	}
}

func (m *FAQTranslationMapper) ToDTO(translation models.FAQTranslation) dto.FAQTranslationResponse {
	return dto.FAQTranslationResponse{
		ID:       translation.ID,
		FAQID:    translation.FAQID,
		Language: translation.LanguageCode,
		Question: translation.Question,
		Answer:   translation.Answer,
	}
}

func (m *FAQTranslationMapper) ToDTOs(translations []models.FAQTranslation) []dto.FAQTranslationResponse {
	dtos := make([]dto.FAQTranslationResponse, len(translations))
	for i, translation := range translations {
		dtos[i] = m.ToDTO(translation)
	}
	return dtos
}
