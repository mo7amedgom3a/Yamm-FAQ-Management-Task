package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
)

type UserService interface {
	GetUser(ctx context.Context, id string) (dto.UserResponse, error)
	// GetAllUsers(ctx context.Context) ([]dto.UserResponse, error) // Repo doesn't have FindAllUsers yet, skipping for now or adding if needed.
	// UpdateUser(ctx context.Context, id string, req dto.UpdateUserRequest) (dto.UserResponse, error) // Need UpdateUserRequest DTO
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	repo   repositories.UserRepository
	mapper *mapper.UserMapper
}

func NewUserService(repo repositories.UserRepository, mapper *mapper.UserMapper) UserService {
	return &userService{
		repo:   repo,
		mapper: mapper,
	}
}

func (s *userService) GetUser(ctx context.Context, id string) (dto.UserResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	user, err := s.repo.FindUserByID(ctx, uid)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return s.mapper.ToDTO(user), nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUser(ctx, uid)
}
