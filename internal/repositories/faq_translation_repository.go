package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"gorm.io/gorm"
)

type FaqTranslationRepository interface {
	CreateFaqTranslation(ctx context.Context, faqTranslation models.FAQTranslation) error
	FindFaqTranslationByID(ctx context.Context, id uuid.UUID) (models.FAQTranslation, error)
	FindFaqTranslationByFaqID(ctx context.Context, faqID uuid.UUID) (models.FAQTranslation, error)
	FindFaqTranslationByLanguage(ctx context.Context, language string) (models.FAQTranslation, error)
	FindByFAQAndLanguage(ctx context.Context, faqID uuid.UUID, language string) (models.FAQTranslation, error)
	UpdateFaqTranslation(ctx context.Context, faqTranslation models.FAQTranslation) error
	DeleteFaqTranslation(ctx context.Context, id uuid.UUID) error
}

type faqTranslationRepository struct {
	db *gorm.DB
}

func NewFaqTranslationRepository(db *gorm.DB) FaqTranslationRepository {
	return &faqTranslationRepository{db: db}
}

func (r *faqTranslationRepository) CreateFaqTranslation(ctx context.Context, faqTranslation models.FAQTranslation) error {
	return r.db.WithContext(ctx).Create(&faqTranslation).Error
}
func (r *faqTranslationRepository) FindFaqTranslationByID(ctx context.Context, id uuid.UUID) (models.FAQTranslation, error) {
	var faqTranslation models.FAQTranslation
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&faqTranslation).Error; err != nil {
		return faqTranslation, err
	}
	return faqTranslation, nil
}

func (r *faqTranslationRepository) FindFaqTranslationByFaqID(ctx context.Context, faqID uuid.UUID) (models.FAQTranslation, error) {
	var faqTranslation models.FAQTranslation
	if err := r.db.WithContext(ctx).Where("faq_id = ?", faqID).First(&faqTranslation).Error; err != nil {
		return faqTranslation, err
	}
	return faqTranslation, nil
}
func (r *faqTranslationRepository) FindFaqTranslationByLanguage(ctx context.Context, language string) (models.FAQTranslation, error) {
	var faqTranslation models.FAQTranslation
	if err := r.db.WithContext(ctx).Where("language = ?", language).First(&faqTranslation).Error; err != nil {
		return faqTranslation, err
	}
	return faqTranslation, nil
}
func (r *faqTranslationRepository) FindByFAQAndLanguage(ctx context.Context, faqID uuid.UUID, language string) (models.FAQTranslation, error) {
	var faqTranslation models.FAQTranslation
	if err := r.db.WithContext(ctx).Where("faq_id = ? AND language = ?", faqID, language).First(&faqTranslation).Error; err != nil {
		return faqTranslation, err
	}
	return faqTranslation, nil
}
func (r *faqTranslationRepository) UpdateFaqTranslation(ctx context.Context, faqTranslation models.FAQTranslation) error {
	return r.db.WithContext(ctx).Save(&faqTranslation).Error
}
func (r *faqTranslationRepository) DeleteFaqTranslation(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.FAQTranslation{}, id).Error
}
