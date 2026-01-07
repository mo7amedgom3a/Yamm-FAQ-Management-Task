package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	FindUserByEmail(ctx context.Context, email string) (models.User, error)
	FindUserByID(ctx context.Context, id uuid.UUID) (models.User, error)
	FindUserByRole(ctx context.Context, role string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindUserByRole(ctx context.Context, role string) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("role = ?", role).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user models.User) error {
	return r.db.WithContext(ctx).Save(&user).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}
