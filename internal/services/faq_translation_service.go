package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

type FAQTranslationService interface {
	CreateTranslation(ctx context.Context, req dto.CreateFAQTranslationRequest) (dto.FAQTranslationResponse, error)
	GetTranslation(ctx context.Context, id string) (dto.FAQTranslationResponse, error)
	GetTranslationsByFAQID(ctx context.Context, faqID string) ([]dto.FAQTranslationResponse, error)
	UpdateTranslation(ctx context.Context, id string, req dto.CreateFAQTranslationRequest) (dto.FAQTranslationResponse, error)
	DeleteTranslation(ctx context.Context, id string) error
}

type faqTranslationService struct {
	repo   repositories.FaqTranslationRepository
	mapper *mapper.FAQTranslationMapper
}

func NewFAQTranslationService(repo repositories.FaqTranslationRepository, mapper *mapper.FAQTranslationMapper) FAQTranslationService {
	return &faqTranslationService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *faqTranslationService) CreateTranslation(ctx context.Context, req dto.CreateFAQTranslationRequest) (dto.FAQTranslationResponse, error) {
	model := s.mapper.ToModel(req)
	model.ID = uuid.New()

	err := s.repo.CreateFaqTranslation(ctx, model)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqTranslationService) GetTranslation(ctx context.Context, id string) (dto.FAQTranslationResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}
	translation, err := s.repo.FindFaqTranslationByID(ctx, uid)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}
	return s.mapper.ToDTO(translation), nil
}

func (s *faqTranslationService) GetTranslationsByFAQID(ctx context.Context, faqID string) ([]dto.FAQTranslationResponse, error) {
	uid, err := uuid.Parse(faqID)
	if err != nil {
		return nil, err
	}
	translations, err := s.repo.FindFaqTranslationByFaqID(ctx, uid)
	if err != nil {
		return nil, err
	}
	return s.mapper.ToDTOs(translations), nil
}

func (s *faqTranslationService) UpdateTranslation(ctx context.Context, id string, req dto.CreateFAQTranslationRequest) (dto.FAQTranslationResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}

	// Check if exists
	translation, err := s.repo.FindFaqTranslationByID(ctx, uid)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}

	model := s.mapper.ToModel(req)
	model.ID = translation.ID
	model.CreatedAt = translation.CreatedAt

	err = s.repo.UpdateFaqTranslation(ctx, model)
	if err != nil {
		return dto.FAQTranslationResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqTranslationService) DeleteTranslation(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteFaqTranslation(ctx, uid)
}
