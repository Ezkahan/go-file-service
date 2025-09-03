package server

import (
	"context"
	"fmt"
	"log"

	"github.com/ezkahan/meditation-backend/internal/config"
	"github.com/ezkahan/meditation-backend/internal/delivery/http"
	"github.com/ezkahan/meditation-backend/internal/repository"
	"github.com/ezkahan/meditation-backend/internal/server/router"
	"github.com/ezkahan/meditation-backend/internal/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

// RunServer initializes DB, handlers, router, and starts the server
func RunHTTPServer() {
	cfg := config.Load()
	// ------------------------
	// Postgres connection
	// ------------------------
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// ------------------------
	// Repositories & Services
	// ------------------------
	categoryRepo := repository.NewCategoryRepo(pool)
	categoryService := usecase.NewCategoryService(categoryRepo)
	categoryHandler := http.NewCategoryHandler(categoryService)

	fileRepo := repository.NewFileRepo(pool)
	fileService := usecase.NewFileService(fileRepo)
	fileHandler := http.NewFileHandler(fileService)

	// ------------------------
	// Router setup
	// ------------------------
	r := router.SetupRouter(categoryHandler, fileHandler)

	// ------------------------
	// Start server
	// ------------------------
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Server running on %s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
