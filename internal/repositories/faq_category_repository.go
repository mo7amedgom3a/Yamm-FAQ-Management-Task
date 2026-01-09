package repositories

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FaqCategoryRepository interface {
	CreateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error
	FindFaqCategoryByID(ctx context.Context, id uuid.UUID) (models.FAQCategory, error)
	FindFaqCategoryByName(ctx context.Context, name string) (models.FAQCategory, error)
	FindAllFaqCategories(ctx context.Context) ([]models.FAQCategory, error)
	UpdateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error
	DeleteFaqCategory(ctx context.Context, id uuid.UUID) error
}

type faqCategoryRepository struct {
	db *gorm.DB
}

func NewFaqCategoryRepository(db *gorm.DB) FaqCategoryRepository {
	return &faqCategoryRepository{db: db}
}

func (r *faqCategoryRepository) CreateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error {
	return r.db.WithContext(ctx).Create(&faqCategory).Error
}
func (r *faqCategoryRepository) FindFaqCategoryByID(ctx context.Context, id uuid.UUID) (models.FAQCategory, error) {
	var faqCategory models.FAQCategory
	if err := r.db.WithContext(ctx).Preload("FAQs").Where("id = ?", id).First(&faqCategory).Error; err != nil {
		return faqCategory, err
	}
	return faqCategory, nil
}
func (r *faqCategoryRepository) FindFaqCategoryByName(ctx context.Context, name string) (models.FAQCategory, error) {
	var faqCategory models.FAQCategory
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&faqCategory).Error; err != nil {
		return faqCategory, err
	}
	return faqCategory, nil
}

func (r *faqCategoryRepository) FindAllFaqCategories(ctx context.Context) ([]models.FAQCategory, error) {
	var faqCategories []models.FAQCategory
	if err := r.db.WithContext(ctx).Preload("FAQs").Find(&faqCategories).Error; err != nil {
		return nil, err
	}
	return faqCategories, nil
}

func (r *faqCategoryRepository) UpdateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error {
	return r.db.WithContext(ctx).Save(&faqCategory).Error
}
func (r *faqCategoryRepository) DeleteFaqCategory(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.FAQCategory{}, id).Error
}
