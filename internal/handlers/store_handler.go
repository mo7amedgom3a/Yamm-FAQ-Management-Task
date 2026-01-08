package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/middleware"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type StoreHandler struct {
	storeService services.StoreService
}

func NewStoreHandler(storeService services.StoreService) *StoreHandler {
	return &StoreHandler{storeService: storeService}
}

func (h *StoreHandler) GetMyStore(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	store, err := h.storeService.GetStoreByMerchantID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func (h *StoreHandler) UpdateStore(c *gin.Context) {
	userID, exists := c.Get(middleware.ContextUserIDKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	store, err := h.storeService.GetStoreByMerchantID(c.Request.Context(), userID.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "store not found"})
		return
	}

	var req dto.CreateStoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	merchantID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		return
	}
	req.MerchantID = merchantID // Enforce ownership

	updatedStore, err := h.storeService.UpdateStore(c.Request.Context(), store.ID.String(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStore)
}
