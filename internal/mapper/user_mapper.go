package mapper

import (
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type UserMapper struct{}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (m *UserMapper) ToModel(req dto.SignupRequest) models.User {
	return models.User{
		Email: req.Email,
		// PasswordHash is handled in service
		Role: models.UserRole(req.Role),
	}
}

func (m *UserMapper) ToDTO(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
	}
}
