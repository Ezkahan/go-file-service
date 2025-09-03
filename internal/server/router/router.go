package router

import (
	"github.com/ezkahan/meditation-backend/internal/delivery/http"
	"github.com/ezkahan/meditation-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter defines all routes and applies middlewares
func SetupRouter(categoryHandler *http.CategoryHandler, fileHandler *http.FileHandler) *gin.Engine {
	r := gin.Default()

	// Health check (public)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Login endpoint (public)
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// TODO: validate credentials
		userID := "123"

		token, err := middleware.GenerateToken(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "could not generate token"})
			return
		}

		c.JSON(200, gin.H{"token": token})
	})

	// Protected routes (JWT middleware)
	protected := r.Group("/")
	protected.Use(middleware.JWTMiddleware())

	categories := protected.Group("/categories")
	{
		categories.POST("", categoryHandler.Create)
		categories.GET("", categoryHandler.List)
		categories.GET("/:id", categoryHandler.Get)
		categories.PUT("/:id", categoryHandler.Update)
		categories.DELETE("/:id", categoryHandler.Delete)
	}

	files := protected.Group("/files")
	{
		files.POST("", fileHandler.Create)
		files.GET("", fileHandler.List)
		files.GET("/:id", fileHandler.Get)
		files.PUT("/:id", fileHandler.Update)
		files.DELETE("/:id", fileHandler.Delete)
	}

	return r
}
