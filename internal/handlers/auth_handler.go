package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/config"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type AuthHandler struct {
	authService services.AuthService
	cfg         *config.Config
}

func NewAuthHandler(authService services.AuthService, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		cfg:         cfg,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate token for the new user
	token, err := h.authService.Login(c.Request.Context(), dto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		// Should not happen if register succeeded, but handle it
		c.JSON(http.StatusOK, gin.H{"user": res, "message": "User registered but auto-login failed", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": res, "token": token})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
