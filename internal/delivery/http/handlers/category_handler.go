package http

import (
	"net/http"

	"github.com/ezkahan/go-file-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *usecase.CategoryService
}

func NewCategoryHandler(s *usecase.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: s}
}

// Create a new category
func (h *CategoryHandler) Create(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Icon     string `json:"icon"`
		ParentId *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	cat, err := h.service.CreateCategory(req.Name, req.Icon, req.ParentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cat)
}

// List all categories
func (h *CategoryHandler) List(c *gin.Context) {
	cats, err := h.service.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cats)
}

// Get a single category by ID
func (h *CategoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	cat, err := h.service.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}
	c.JSON(http.StatusOK, cat)
}

// Update category
func (h *CategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name     string `json:"name" binding:"required"`
		Icon     string `json:"icon"`
		ParentId *uint  `json:"parent_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.service.UpdateCategory(id, req.Name, req.Icon, req.ParentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// Delete category
func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
