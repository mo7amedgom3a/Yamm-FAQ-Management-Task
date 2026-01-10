package servicetests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) CreateStore(ctx context.Context, store models.Store) error {
	args := m.Called(ctx, store)
	return args.Error(0)
}

func (m *MockStoreRepository) FindStoreByID(ctx context.Context, id uuid.UUID) (models.Store, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Store), args.Error(1)
}

func (m *MockStoreRepository) FindStoreByName(ctx context.Context, name string) (models.Store, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.Store), args.Error(1)
}

func (m *MockStoreRepository) FindStoreByMerchantID(ctx context.Context, merchantID uuid.UUID) (models.Store, error) {
	args := m.Called(ctx, merchantID)
	return args.Get(0).(models.Store), args.Error(1)
}

func (m *MockStoreRepository) UpdateStore(ctx context.Context, store models.Store) error {
	args := m.Called(ctx, store)
	return args.Error(0)
}

func (m *MockStoreRepository) DeleteStore(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type StoreServiceSuite struct {
	suite.Suite
	repo    *MockStoreRepository
	service services.StoreService
}

func (s *StoreServiceSuite) SetupTest() {
	s.repo = new(MockStoreRepository)
	transMapper := mapper.NewFAQTranslationMapper()
	faqMapper := mapper.NewFAQMapper(transMapper)
	storeMapper := mapper.NewStoreMapper(faqMapper)
	s.service = services.NewStoreService(s.repo, storeMapper)
}

func (s *StoreServiceSuite) TestGetStore() {
	id := uuid.New()
	store := models.Store{
		ID:   id,
		Name: "My Store",
	}

	s.repo.On("FindStoreByID", mock.Anything, id).Return(store, nil)

	resp, err := s.service.GetStore(context.Background(), id.String())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "My Store", resp.Name)
	s.repo.AssertExpectations(s.T())
}

func TestStoreServiceSuite(t *testing.T) {
	suite.Run(t, new(StoreServiceSuite))
}
