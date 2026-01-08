package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/middleware"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	// Admin can delete any user, or user can delete themselves?
	// For now, let's implement delete self or admin delete by ID.
	// The route param will decide.
	id := c.Param("id")
	if id == "" {
		// If no ID, delete self?
		userID, _ := c.Get(middleware.ContextUserIDKey)
		id = userID.(string)
	}

	err := h.userService.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
