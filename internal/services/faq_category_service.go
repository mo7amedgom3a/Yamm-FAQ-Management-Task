package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

type FAQCategoryService interface {
	CreateCategory(ctx context.Context, req dto.CreateFAQCategoryRequest) (dto.FAQCategoryResponse, error)
	GetCategory(ctx context.Context, id string) (dto.FAQCategoryResponse, error)
	GetAllCategories(ctx context.Context) ([]dto.FAQCategoryResponse, error)
	UpdateCategory(ctx context.Context, id string, req dto.CreateFAQCategoryRequest) (dto.FAQCategoryResponse, error)
	DeleteCategory(ctx context.Context, id string) error
}

type faqCategoryService struct {
	repo   repositories.FaqCategoryRepository
	mapper *mapper.FAQCategoryMapper
}

func NewFAQCategoryService(repo repositories.FaqCategoryRepository, mapper *mapper.FAQCategoryMapper) FAQCategoryService {
	return &faqCategoryService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *faqCategoryService) CreateCategory(ctx context.Context, req dto.CreateFAQCategoryRequest) (dto.FAQCategoryResponse, error) {
	model := s.mapper.ToModel(req)
	model.ID = uuid.New() // Generate ID here or in repo? Usually here or DB. Let's do here.

	err := s.repo.CreateFaqCategory(ctx, model)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqCategoryService) GetCategory(ctx context.Context, id string) (dto.FAQCategoryResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}
	category, err := s.repo.FindFaqCategoryByID(ctx, uid)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}
	return s.mapper.ToDTO(category), nil
}

func (s *faqCategoryService) GetAllCategories(ctx context.Context) ([]dto.FAQCategoryResponse, error) {
	categories, err := s.repo.FindAllFaqCategories(ctx)
	if err != nil {
		return nil, err
	}
	return s.mapper.ToDTOs(categories), nil
}

func (s *faqCategoryService) UpdateCategory(ctx context.Context, id string, req dto.CreateFAQCategoryRequest) (dto.FAQCategoryResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}

	// Check if exists
	category, err := s.repo.FindFaqCategoryByID(ctx, uid)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}

	model := s.mapper.ToModel(req)
	model.ID = category.ID               // Keep ID
	model.CreatedAt = category.CreatedAt // Keep CreatedAt

	err = s.repo.UpdateFaqCategory(ctx, model)
	if err != nil {
		return dto.FAQCategoryResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}

func (s *faqCategoryService) DeleteCategory(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteFaqCategory(ctx, uid)
}
