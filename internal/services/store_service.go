package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

type StoreService interface {
	GetStore(ctx context.Context, id string) (dto.StoreResponse, error)
	GetStoreByMerchantID(ctx context.Context, merchantID string) (dto.StoreResponse, error)
	UpdateStore(ctx context.Context, id string, req dto.CreateStoreRequest) (dto.StoreResponse, error)
}

type storeService struct {
	repo   repositories.StoreRepository
	mapper *mapper.StoreMapper
}

func NewStoreService(repo repositories.StoreRepository, mapper *mapper.StoreMapper) StoreService {
	return &storeService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *storeService) GetStore(ctx context.Context, id string) (dto.StoreResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	store, err := s.repo.FindStoreByID(ctx, uid)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	return s.mapper.ToDTO(store), nil
}

func (s *storeService) GetStoreByMerchantID(ctx context.Context, merchantID string) (dto.StoreResponse, error) {
	uid, err := uuid.Parse(merchantID)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	store, err := s.repo.FindStoreByMerchantID(ctx, uid)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	return s.mapper.ToDTO(store), nil
}

func (s *storeService) UpdateStore(ctx context.Context, id string, req dto.CreateStoreRequest) (dto.StoreResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.StoreResponse{}, err
	}

	store, err := s.repo.FindStoreByID(ctx, uid)
	if err != nil {
		return dto.StoreResponse{}, err
	}

	model := s.mapper.ToModel(req)
	model.ID = store.ID
	model.CreatedAt = store.CreatedAt

	err = s.repo.UpdateStore(ctx, model)
	if err != nil {
		return dto.StoreResponse{}, err
	}
	return s.mapper.ToDTO(model), nil
}
