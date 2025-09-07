include .env
export $(shell sed 's/=.*//' .env)

# Variables
APP_NAME := go-file-service
ENV_FILE := .env
BIN := bin/$(APP_NAME)
DB_MIGRATE := migrate # assuming you use a migration tool like golang-migrate
MIGRATION_PATH := internal/db/migrations

# Default target
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make run           # Run server locally"
	@echo "  make build         # Build binary"
	@echo "  make migrate-up    # Run DB migrations up"
	@echo "  make migrate-down  # Rollback last DB migration"
	@echo "  make deploy        # Deploy application"

# Run server locally
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	go run cmd/server/main.go
# 	@export $(shell cat $(ENV_FILE) | xargs) && go run cmd/server/main.go

# Build binary
.PHONY: build
build:
	@echo "Building $(APP_NAME) binary..."
	@mkdir -p bin
	@go build -o $(BIN) cmd/server/main.go
	@echo "Built binary: $(BIN)"

# Run DB migrations up
.PHONY: migrate-up
migrate-up:
	@echo "Running migrations up... $$DATABASE_URL"
	@$(DB_MIGRATE) -path $(MIGRATION_PATH) -database $$DATABASE_URL up

# Rollback last migration
.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last migration..."
	@$(DB_MIGRATE) -path $(MIGRATION_PATH) -database $$DATABASE_URL down 1

# Deploy application
.PHONY: deploy
deploy: build
	@echo "Deploying $(APP_NAME)..."
	# Example deployment steps:
	# Copy binary to server, restart service, etc.
	@scp $(BIN) user@server:/path/to/deploy/
	@ssh user@server 'systemctl restart $(APP_NAME)'
	@echo "Deployment completed!"
