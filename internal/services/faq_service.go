package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

type FAQService interface {
	CreateFAQ(ctx context.Context, req dto.CreateFAQRequest) (dto.FAQResponse, error)
	GetFAQ(ctx context.Context, id string) (dto.FAQResponse, error)
	GetAllFAQs(ctx context.Context, storeID string) ([]dto.FAQResponse, error)
	UpdateFAQ(ctx context.Context, id string, req dto.CreateFAQRequest) (dto.FAQResponse, error)
	DeleteFAQ(ctx context.Context, id string) error
}

type faqService struct {
	repo   repositories.FaqRepository
	mapper *mapper.FAQMapper
}

func NewFAQService(repo repositories.FaqRepository, mapper *mapper.FAQMapper) FAQService {
	return &faqService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *faqService) CreateFAQ(ctx context.Context, req dto.CreateFAQRequest) (dto.FAQResponse, error) {
	model := s.mapper.ToModel(req)
	model.ID = uuid.New()

	err := s.repo.CreateFaq(ctx, model)
	if err != nil {
		return dto.FAQResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqService) GetFAQ(ctx context.Context, id string) (dto.FAQResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQResponse{}, err
	}
	faq, err := s.repo.FindFaqByID(ctx, uid)
	if err != nil {
		return dto.FAQResponse{}, err
	}
	return s.mapper.ToDTO(faq), nil
}

func (s *faqService) GetAllFAQs(ctx context.Context, storeID string) ([]dto.FAQResponse, error) {
	var uid uuid.UUID
	var err error

	if storeID != "" {
		uid, err = uuid.Parse(storeID)
		if err != nil {
			return nil, err
		}
	} else {
		uid = uuid.Nil
	}

	faqs, err := s.repo.FindGlobalAndStoreFAQs(ctx, uid)
	if err != nil {
		return nil, err
	}
	return s.mapper.ToDTOs(faqs), nil
}

func (s *faqService) UpdateFAQ(ctx context.Context, id string, req dto.CreateFAQRequest) (dto.FAQResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQResponse{}, err
	}

	// Check if exists
	faq, err := s.repo.FindFaqByID(ctx, uid)
	if err != nil {
		return dto.FAQResponse{}, err
	}

	model := s.mapper.ToModel(req)
	model.ID = faq.ID
	model.CreatedAt = faq.CreatedAt

	err = s.repo.UpdateFaq(ctx, model)
	if err != nil {
		return dto.FAQResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqService) DeleteFAQ(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteFaq(ctx, uid)
}
