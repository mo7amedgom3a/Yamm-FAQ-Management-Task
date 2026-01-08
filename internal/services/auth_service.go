package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/auth"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/mapper"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, req dto.SignupRequest) (dto.UserResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
	VerifyToken(ctx context.Context, token string) (dto.UserResponse, error)
}

type authService struct {
	userRepo    repositories.UserRepository
	storeRepo   repositories.StoreRepository
	cfg         *config.Config
	userMapper  *mapper.UserMapper
	storeMapper *mapper.StoreMapper
}

func NewAuthService(userRepo repositories.UserRepository, storeRepo repositories.StoreRepository, cfg *config.Config, userMapper *mapper.UserMapper, storeMapper *mapper.StoreMapper) AuthService {
	return &authService{
		userRepo:    userRepo,
		storeRepo:   storeRepo,
		cfg:         cfg,
		userMapper:  userMapper,
		storeMapper: storeMapper,
	}
}

func (s *authService) Register(ctx context.Context, req dto.SignupRequest) (dto.UserResponse, error) {
	// Check if user exists
	_, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	if err == nil {
		return dto.UserResponse{}, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserResponse{}, err
	}

	user := s.userMapper.ToModel(req)
	user.ID = uuid.New()
	user.PasswordHash = string(hashedPassword)

	// Set default role if empty
	if user.Role == "" {
		user.Role = "user"
	}

	err = s.userRepo.CreateUser(ctx, &user)
	if err != nil {
		return dto.UserResponse{}, err
	}
	// create default store when user created if role is merchant
	if user.Role == "merchant" {
		storeReq := dto.CreateStoreRequest{
			Name:       "Default Store",
			MerchantID: user.ID,
		}
		store := s.storeMapper.ToModel(storeReq)
		store.ID = uuid.New()

		err = s.storeRepo.CreateStore(ctx, store)
		if err != nil {
			return dto.UserResponse{}, err
		}
	}

	return s.userMapper.ToDTO(user), nil
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := auth.GenerateToken(&user, s.cfg)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) VerifyToken(ctx context.Context, token string) (dto.UserResponse, error) {
	err := auth.VerifyToken(token, s.cfg)
	if err != nil {
		return dto.UserResponse{}, err
	}

	// TODO: Extract claims and get user
	claims, err := auth.ExtractClaims(token, s.cfg)
	if err != nil {
		return dto.UserResponse{}, err
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return dto.UserResponse{}, errors.New("invalid user_id in token")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return dto.UserResponse{}, errors.New("invalid user_id format")
	}

	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return s.userMapper.ToDTO(user), nil
}
