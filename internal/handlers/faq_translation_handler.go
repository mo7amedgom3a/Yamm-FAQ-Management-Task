package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type FAQTranslationHandler struct {
	translationService services.FAQTranslationService
}

func NewFAQTranslationHandler(translationService services.FAQTranslationService) *FAQTranslationHandler {
	return &FAQTranslationHandler{translationService: translationService}
}

func (h *FAQTranslationHandler) CreateTranslation(c *gin.Context) {
	var req dto.CreateFAQTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.translationService.CreateTranslation(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *FAQTranslationHandler) GetTranslation(c *gin.Context) {
	id := c.Param("id")
	res, err := h.translationService.GetTranslation(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "translation not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQTranslationHandler) GetTranslationsByFAQID(c *gin.Context) {
	faqID := c.Param("id")
	res, err := h.translationService.GetTranslationsByFAQID(c.Request.Context(), faqID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQTranslationHandler) UpdateTranslation(c *gin.Context) {
	id := c.Param("id")
	var req dto.CreateFAQTranslationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.translationService.UpdateTranslation(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQTranslationHandler) DeleteTranslation(c *gin.Context) {
	id := c.Param("id")
	err := h.translationService.DeleteTranslation(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "translation deleted"})
}
