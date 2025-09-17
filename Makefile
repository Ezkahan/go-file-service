include .env
export $(shell sed 's/=.*//' .env)

# Variables
APP_NAME := meditation
BIN_PATH := bin/$(APP_NAME)
ENV_FILE := .env
USER := ubuntu
HOST := 216.250.11.236
SSH_PORT := 22020
DEPLOY_PATH := /var/www/apps/meditation
SERVICE_NAME := meditation
DB_MIGRATE := migrate
MIGRATION_PATH := internal/db/migrations

# Default target
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Usage:"
	@echo "  make run           # Run server locally"
	@echo "  make build         # Build binary"
	@echo "  make optimize      # Compress binary with upx"
	@echo "  make migrate-up    # Run DB migrations up"
	@echo "  make migrate-down  # Rollback last DB migration"
	@echo "  make deploy        # Deploy application"

# Run server locally
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	go run cmd/app/main.go

# Build binary
.PHONY: build
build:
	@echo "Building $(APP_NAME) binary..."
	@mkdir -p bin
	GOOS=linux GOARCH=amd64 go build \
		-ldflags="-s -w -X main.Version=1.0.0 -X main.BuildTime=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)" \
		-o $(BIN_PATH) ./cmd/app/main.go
	@echo "Built binary: $(BIN_PATH)"

# Optimize binary
.PHONY: optimize
optimize: build
	@echo "Optimizing $(BIN_PATH) with upx..."
	upx --best --lzma $(BIN_PATH)

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

# Deploy application (build + optimize + scp + restart service)
.PHONY: deploy
deploy: optimize
	@echo "Deploying $(APP_NAME)..."
	scp -P $(SSH_PORT) $(BIN_PATH) $(USER)@$(HOST):$(DEPLOY_PATH)
	ssh -p $(SSH_PORT) $(USER)@$(HOST) 'sudo systemctl restart $(SERVICE_NAME)'
	@echo "Deployment completed!"
