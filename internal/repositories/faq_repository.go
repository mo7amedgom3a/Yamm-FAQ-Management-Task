package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"context"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type FaqRepository interface {
	CreateFaq(ctx context.Context, faq models.FAQ) error
	FindFaqByID(ctx context.Context, id uuid.UUID) (models.FAQ, error)
	FindFaqByName(ctx context.Context, name string) (models.FAQ, error)
	FindFaqByStoreID(ctx context.Context, storeID uuid.UUID) (models.FAQ, error)
	FindFaqByCategoryID(ctx context.Context, categoryID uuid.UUID) (models.FAQ, error)
	FindGlobalAndStoreFAQs(ctx context.Context, storeID uuid.UUID) ([]models.FAQ, error)
	UpdateFaq(ctx context.Context, faq models.FAQ) error
	DeleteFaq(ctx context.Context, id uuid.UUID) error
}

type faqRepository struct {
	db *gorm.DB
}

func NewFaqRepository(db *gorm.DB) FaqRepository {
	return &faqRepository{db: db}
}

func (r *faqRepository) CreateFaq(ctx context.Context, faq models.FAQ) error {
	return r.db.WithContext(ctx).Create(&faq).Error
}
func (r *faqRepository) FindFaqByID(ctx context.Context, id uuid.UUID) (models.FAQ, error) {
	var faq models.FAQ
	if err := r.db.WithContext(ctx).Preload("Translations").Where("id = ?", id).First(&faq).Error; err != nil {
		return faq, err
	}
	return faq, nil
}
func (r *faqRepository) FindFaqByName(ctx context.Context, name string) (models.FAQ, error) {
	var faq models.FAQ
	if err := r.db.WithContext(ctx).Preload("Translations").Where("name = ?", name).First(&faq).Error; err != nil {
		return faq, err
	}
	return faq, nil
}
func (r *faqRepository) FindFaqByStoreID(ctx context.Context, storeID uuid.UUID) (models.FAQ, error) {
	var faq models.FAQ
	if err := r.db.WithContext(ctx).Preload("Translations").Where("store_id = ?", storeID).First(&faq).Error; err != nil {
		return faq, err
	}
	return faq, nil
}
func (r *faqRepository) FindFaqByCategoryID(ctx context.Context, categoryID uuid.UUID) (models.FAQ, error) {
	var faq models.FAQ
	if err := r.db.WithContext(ctx).Preload("Translations").Where("category_id = ?", categoryID).First(&faq).Error; err != nil {
		return faq, err
	}
	return faq, nil
}
func (r *faqRepository) FindGlobalAndStoreFAQs(ctx context.Context, storeID uuid.UUID) ([]models.FAQ, error) {
	var faqs []models.FAQ
	err := r.db.
		Preload("Translations").
		Where("is_global = TRUE OR store_id = ?", storeID).
		Find(&faqs).
		Error
	if err != nil {
		return nil, err
	}
	return faqs, nil
}
func (r *faqRepository) UpdateFaq(ctx context.Context, faq models.FAQ) error {
	return r.db.WithContext(ctx).Save(&faq).Error
}
func (r *faqRepository) DeleteFaq(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.FAQ{}, id).Error
}
