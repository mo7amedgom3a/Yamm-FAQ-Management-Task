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

type MockFaqCategoryRepository struct {
	mock.Mock
}

func (m *MockFaqCategoryRepository) CreateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error {
	args := m.Called(ctx, faqCategory)
	return args.Error(0)
}

func (m *MockFaqCategoryRepository) FindFaqCategoryByID(ctx context.Context, id uuid.UUID) (models.FAQCategory, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.FAQCategory), args.Error(1)
}

func (m *MockFaqCategoryRepository) FindFaqCategoryByName(ctx context.Context, name string) (models.FAQCategory, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(models.FAQCategory), args.Error(1)
}

func (m *MockFaqCategoryRepository) FindAllFaqCategories(ctx context.Context) ([]models.FAQCategory, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.FAQCategory), args.Error(1)
}

func (m *MockFaqCategoryRepository) UpdateFaqCategory(ctx context.Context, faqCategory models.FAQCategory) error {
	args := m.Called(ctx, faqCategory)
	return args.Error(0)
}

func (m *MockFaqCategoryRepository) DeleteFaqCategory(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type FAQCategoryServiceSuite struct {
	suite.Suite
	repo    *MockFaqCategoryRepository
	service services.FAQCategoryService
}

func (s *FAQCategoryServiceSuite) SetupTest() {
	s.repo = new(MockFaqCategoryRepository)
	transMapper := mapper.NewFAQTranslationMapper()
	faqMapper := mapper.NewFAQMapper(transMapper)
	catMapper := mapper.NewFAQCategoryMapper(faqMapper)
	s.service = services.NewFAQCategoryService(s.repo, catMapper)
}

func (s *FAQCategoryServiceSuite) TestCreateCategory() {
	req := dto.CreateFAQCategoryRequest{
		Name: "General",
	}

	s.repo.On("CreateFaqCategory", mock.Anything, mock.MatchedBy(func(cat models.FAQCategory) bool {
		return cat.Name == req.Name
	})).Return(nil)

	resp, err := s.service.CreateCategory(context.Background(), req)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), req.Name, resp.Name)
	s.repo.AssertExpectations(s.T())
}

func (s *FAQCategoryServiceSuite) TestGetAllCategories() {
	categories := []models.FAQCategory{
		{Name: "Cat1"},
		{Name: "Cat2"},
	}

	s.repo.On("FindAllFaqCategories", mock.Anything).Return(categories, nil)

	resp, err := s.service.GetAllCategories(context.Background())
	assert.NoError(s.T(), err)
	assert.Len(s.T(), resp, 2)
	assert.Equal(s.T(), "Cat1", resp[0].Name)
	s.repo.AssertExpectations(s.T())
}

func TestFAQCategoryServiceSuite(t *testing.T) {
	suite.Run(t, new(FAQCategoryServiceSuite))
}
