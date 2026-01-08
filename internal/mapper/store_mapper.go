package mapper

import (
	"time"

	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
)

type StoreMapper struct {
	faqMapper *FAQMapper
}

func NewStoreMapper(faqMapper *FAQMapper) *StoreMapper {
	return &StoreMapper{
		faqMapper: faqMapper,
	}
}

func (m *StoreMapper) ToModel(req dto.CreateStoreRequest) models.Store {
	return models.Store{
		Name:       req.Name,
		MerchantID: req.MerchantID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (m *StoreMapper) ToDTO(store models.Store) dto.StoreResponse {
	return dto.StoreResponse{
		ID:         store.ID,
		Name:       store.Name,
		MerchantID: store.MerchantID,
		FAQs:       m.faqMapper.ToDTOs(store.FAQs),
	}
}

func (m *StoreMapper) ToDTOs(stores []models.Store) []dto.StoreResponse {
	dtos := make([]dto.StoreResponse, len(stores))
	for i, store := range stores {
		dtos[i] = m.ToDTO(store)
	}
	return dtos
}
