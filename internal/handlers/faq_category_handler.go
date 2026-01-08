package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/dto"
	"github.com/mo7amedgom3a/Yamm-FAQ-Management-Task/internal/services"
)

type FAQCategoryHandler struct {
	categoryService services.FAQCategoryService
}

func NewFAQCategoryHandler(categoryService services.FAQCategoryService) *FAQCategoryHandler {
	return &FAQCategoryHandler{categoryService: categoryService}
}

func (h *FAQCategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateFAQCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.categoryService.CreateCategory(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *FAQCategoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	res, err := h.categoryService.GetCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQCategoryHandler) GetAllCategories(c *gin.Context) {
	res, err := h.categoryService.GetAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQCategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req dto.CreateFAQCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.categoryService.UpdateCategory(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *FAQCategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := h.categoryService.DeleteCategory(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted"})
}
