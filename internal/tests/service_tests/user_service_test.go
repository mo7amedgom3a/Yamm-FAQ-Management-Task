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

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) FindUserByRole(ctx context.Context, role string) (models.User, error) {
	args := m.Called(ctx, role)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

type UserServiceSuite struct {
	suite.Suite
	repo    *MockUserRepository
	service services.UserService
}

func (s *UserServiceSuite) SetupTest() {
	s.repo = new(MockUserRepository)
	mapper := mapper.NewUserMapper()
	s.service = services.NewUserService(s.repo, mapper)
}

func (s *UserServiceSuite) TestGetUser() {
	id := uuid.New()
	user := models.User{
		ID:    id,
		Email: "test@example.com",
		Role:  models.RoleCustomer,
	}

	s.repo.On("FindUserByID", mock.Anything, id).Return(user, nil)

	resp, err := s.service.GetUser(context.Background(), id.String())
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), user.Email, resp.Email)
	s.repo.AssertExpectations(s.T())
}

func (s *UserServiceSuite) TestDeleteUser() {
	id := uuid.New()

	s.repo.On("DeleteUser", mock.Anything, id).Return(nil)

	err := s.service.DeleteUser(context.Background(), id.String())
	assert.NoError(s.T(), err)
	s.repo.AssertExpectations(s.T())
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}
