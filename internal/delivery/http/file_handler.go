package http

import (
	"net/http"

	"github.com/ezkahan/meditation-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	service *usecase.FileService
}

func NewFileHandler(s *usecase.FileService) *FileHandler {
	return &FileHandler{service: s}
}

// Create a new file
func (h *FileHandler) Create(c *gin.Context) {
	var req struct {
		Name       string  `json:"name" binding:"required"`
		IconPath   string  `json:"icon_path"`
		FilePath   string  `json:"file_path" binding:"required"`
		CategoryID *string `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	f, err := h.service.CreateFile(req.Name, req.IconPath, req.FilePath, req.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, f)
}

// List all files
func (h *FileHandler) List(c *gin.Context) {
	files, err := h.service.ListFiles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, files)
}

// Get a file by ID
func (h *FileHandler) Get(c *gin.Context) {
	id := c.Param("id")
	f, err := h.service.GetFile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	c.JSON(http.StatusOK, f)
}

// Update a file
func (h *FileHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Name       string  `json:"name" binding:"required"`
		IconPath   string  `json:"icon_path"`
		FilePath   string  `json:"file_path" binding:"required"`
		CategoryID *string `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.service.UpdateFile(id, req.Name, req.IconPath, req.FilePath, req.CategoryID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// Delete a file
func (h *FileHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteFile(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
