package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(ctx context.Context, store models.Store) error
	FindStoreByID(ctx context.Context, id uuid.UUID) (models.Store, error)
	FindStoreByName(ctx context.Context, name string) (models.Store, error)
	FindStoreByMerchantID(ctx context.Context, merchantID uuid.UUID) (models.Store, error)
	UpdateStore(ctx context.Context, store models.Store) error
	DeleteStore(ctx context.Context, id uuid.UUID) error
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) CreateStore(ctx context.Context, store models.Store) error {
	return r.db.WithContext(ctx).Create(&store).Error
}
func (r *storeRepository) FindStoreByID(ctx context.Context, id uuid.UUID) (models.Store, error) {
	var store models.Store
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&store).Error; err != nil {
		return store, err
	}
	return store, nil
}
func (r *storeRepository) FindStoreByName(ctx context.Context, name string) (models.Store, error) {
	var store models.Store
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&store).Error; err != nil {
		return store, err
	}
	return store, nil
}
func (r *storeRepository) FindStoreByMerchantID(ctx context.Context, merchantID uuid.UUID) (models.Store, error) {
	var store models.Store
	if err := r.db.WithContext(ctx).Where("merchant_id = ?", merchantID).First(&store).Error; err != nil {
		return store, err
	}
	return store, nil
}

func (r *storeRepository) UpdateStore(ctx context.Context, store models.Store) error {
	return r.db.WithContext(ctx).Save(&store).Error
}
func (r *storeRepository) DeleteStore(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Store{}, id).Error
}
