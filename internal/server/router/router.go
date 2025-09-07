package router

import (
	httpHandler "github.com/ezkahan/go-file-service/internal/delivery/http/handlers"
	"github.com/ezkahan/go-file-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(categoryHandler *httpHandler.CategoryHandler, fileHandler httpHandler.FileHandler, authHandler httpHandler.AuthHandler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api.POST("/login", authHandler.Login)

	protected := api.Group("/admin")
	protected.Use(middleware.JWTAuthMiddleware())

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

	// public routes
	api.GET("/categories", categoryHandler.List)
	api.GET("/categories/:id", categoryHandler.Get)
	api.GET("/files", fileHandler.List)
	api.GET("/files/:id", fileHandler.Get)

	return r
}
