package servicetests

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/models"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockFaqRepository struct {
	mock.Mock
}

func (m *MockFaqRepository) CreateFaq(ctx context.Context, faq models.FAQ) error {
	args := m.Called(ctx, faq)
	return args.Error(0)
}

func (m *MockFaqRepository) FindFaqByID(ctx context.Context, id uuid.UUID) (models.FAQ, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) FindFaqByName(ctx context.Context, name string) (models.FAQ, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) FindFaqByStoreID(ctx context.Context, storeID uuid.UUID) (models.FAQ, error) {
	args := m.Called(ctx, storeID)
	return args.Get(0).(models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) FindFaqsByStoreID(ctx context.Context, storeID uuid.UUID) ([]models.FAQ, error) {
	args := m.Called(ctx, storeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) FindFaqByCategoryID(ctx context.Context, categoryID uuid.UUID) (models.FAQ, error) {
	args := m.Called(ctx, categoryID)
	return args.Get(0).(models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) FindGlobalAndStoreFAQs(ctx context.Context, storeID uuid.UUID) ([]models.FAQ, error) {
	args := m.Called(ctx, storeID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.FAQ), args.Error(1)
}

func (m *MockFaqRepository) UpdateFaq(ctx context.Context, faq models.FAQ) error {
	args := m.Called(ctx, faq)
	return args.Error(0)
}

func (m *MockFaqRepository) DeleteFaq(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type FAQServiceSuite struct {
	suite.Suite
	repo    *MockFaqRepository
	service services.FAQService
}

func (s *FAQServiceSuite) SetupTest() {
	s.repo = new(MockFaqRepository)
	transMapper := mapper.NewFAQTranslationMapper()
	mapper := mapper.NewFAQMapper(transMapper)
	s.service = services.NewFAQService(s.repo, mapper)
}

func (s *FAQServiceSuite) TestCreateFAQ() {
	catID := uuid.New()
	req := dto.CreateFAQRequest{
		CategoryID: catID,
		IsGlobal:   true,
	}

	s.repo.On("CreateFaq", mock.Anything, mock.MatchedBy(func(faq models.FAQ) bool {
		return faq.IsGlobal == true && faq.CategoryID == catID
	})).Return(nil)

	resp, err := s.service.CreateFAQ(context.Background(), req)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), req.IsGlobal, resp.IsGlobal)
	assert.Equal(s.T(), req.CategoryID, resp.CategoryID)
	s.repo.AssertExpectations(s.T())
}

func (s *FAQServiceSuite) TestGetFAQ() {
	id := uuid.New()
	faq := models.FAQ{
		ID: id,
		Translations: []models.FAQTranslation{
			{
				LanguageCode: "en",
				Question:     "Q",
				Answer:       "A",
			},
		},
	}

	s.repo.On("FindFaqByID", mock.Anything, id).Return(faq, nil)

	resp, err := s.service.GetFAQ(context.Background(), id.String())
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), resp.Translations)
	assert.Equal(s.T(), "Q", resp.Translations[0].Question)
	s.repo.AssertExpectations(s.T())
}

func TestFAQServiceSuite(t *testing.T) {
	suite.Run(t, new(FAQServiceSuite))
}
