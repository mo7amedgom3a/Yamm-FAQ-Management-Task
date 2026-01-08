package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/middleware"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type FAQHandler struct {
	faqService   services.FAQService
	storeService services.StoreService
}

func NewFAQHandler(faqService services.FAQService, storeService services.StoreService) *FAQHandler {
	return &FAQHandler{
		faqService:   faqService,
		storeService: storeService,
	}
}

func (h *FAQHandler) CreateFAQ(c *gin.Context) {
	var req dto.CreateFAQRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get(middleware.ContextUserIDKey)
	req.CreatedBy = userID.(string)

	role, _ := c.Get(middleware.ContextUserRoleKey)
	if role == "merchant" {
		store, err := h.storeService.GetStoreByMerchantID(c.Request.Context(), userID.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Merchant must have a store to create FAQs"})
			return
		}
		req.StoreID = &store.ID
		req.IsGlobal = false
	} else if role == "admin" {
		req.IsGlobal = true
		req.StoreID = nil
	}

	res, err := h.faqService.CreateFAQ(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)

}

func (h *FAQHandler) GetFAQ(c *gin.Context) {
	id := c.Param("id")
	res, err := h.faqService.GetFAQ(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQHandler) GetAllFAQs(c *gin.Context) {
	storeID := c.Query("store_id")
	res, err := h.faqService.GetAllFAQs(c.Request.Context(), storeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQHandler) UpdateFAQ(c *gin.Context) {
	id := c.Param("id")
	var req dto.CreateFAQRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.faqService.UpdateFAQ(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQHandler) DeleteFAQ(c *gin.Context) {
	id := c.Param("id")
	err := h.faqService.DeleteFAQ(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "FAQ deleted"})
}
