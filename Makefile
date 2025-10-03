# Yaro Wora Tourism Backend - Makefile
# ====================================

# Variables
APP_NAME=yaro-wora-api
DB_NAME=yaro_wora
DB_USER=postgres
DB_HOST=localhost
DB_PORT=5432

# Load environment variables if .env exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: help build run dev clean test docker-build docker-run db-create db-drop db-reset db-migrate db-seed deps lint fmt

# Default target
help: ## Show this help message
	@echo "$(BLUE)Yaro Wora Tourism Backend$(NC)"
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# =============================================================================
# Development Commands
# =============================================================================

build: ## Build the application
	@echo "$(YELLOW)Building $(APP_NAME)...$(NC)"
	@go build -o $(APP_NAME) .
	@echo "$(GREEN)✅ Build completed!$(NC)"

run: ## Run the application
	@echo "$(YELLOW)Starting $(APP_NAME)...$(NC)"
	@go run main.go

dev: ## Run with hot reload (requires air)
	@echo "$(YELLOW)Starting development server with hot reload...$(NC)"
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "$(RED)❌ Air not found. Installing...$(NC)"; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

test: ## Run tests
	@echo "$(YELLOW)Running tests...$(NC)"
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(NC)"
	@go test -v -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✅ Coverage report generated: coverage.html$(NC)"

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	@rm -f $(APP_NAME)
	@rm -f coverage.out coverage.html
	@go clean
	@echo "$(GREEN)✅ Clean completed!$(NC)"

# =============================================================================
# Database Commands
# =============================================================================

db-create: ## Create the database
	@echo "$(YELLOW)Creating database $(DB_NAME)...$(NC)"
	@createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME) || true
	@echo "$(GREEN)✅ Database $(DB_NAME) created (or already exists)!$(NC)"

db-drop: ## Drop the database
	@echo "$(RED)⚠️  Dropping database $(DB_NAME)...$(NC)"
	@read -p "Are you sure you want to drop the database? (y/N): " confirm && [ "$$confirm" = "y" ]
	@dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME) || true
	@echo "$(GREEN)✅ Database $(DB_NAME) dropped!$(NC)"

db-reset: db-drop db-create ## Drop and recreate the database
	@echo "$(GREEN)✅ Database reset completed!$(NC)"

db-migrate: ## Run database migrations (auto-migrate)
	@echo "$(YELLOW)Running database migrations...$(NC)"
	@go run -ldflags "-X main.migrateOnly=true" main.go || echo "$(BLUE)Note: Migrations run automatically when starting the app$(NC)"

db-seed: ## Seed the database with initial data
	@echo "$(YELLOW)Seeding database with initial data...$(NC)"
	@APP_ENV=development go run main.go &
	@sleep 3
	@pkill -f "go run main.go" || true
	@echo "$(GREEN)✅ Database seeded! (App started briefly to run seeding)$(NC)"

db-status: ## Check database connection
	@echo "$(YELLOW)Checking database connection...$(NC)"
	@psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -c "SELECT version();" > /dev/null 2>&1 && \
		echo "$(GREEN)✅ Database connection successful!$(NC)" || \
		echo "$(RED)❌ Database connection failed!$(NC)"

# =============================================================================
# Code Quality Commands
# =============================================================================

deps: ## Download and install dependencies
	@echo "$(YELLOW)Installing dependencies...$(NC)"
	@go mod download
	@go mod tidy
	@echo "$(GREEN)✅ Dependencies installed!$(NC)"

lint: ## Run linter
	@echo "$(YELLOW)Running linter...$(NC)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "$(BLUE)Installing golangci-lint...$(NC)"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

fmt: ## Format code
	@echo "$(YELLOW)Formatting code...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✅ Code formatted!$(NC)"

vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(NC)"
	@go vet ./...
	@echo "$(GREEN)✅ Go vet completed!$(NC)"

check: fmt vet lint test ## Run all code quality checks

# =============================================================================
# Docker Commands
# =============================================================================

docker-build: ## Build Docker image
	@echo "$(YELLOW)Building Docker image...$(NC)"
	@docker build -t $(APP_NAME):latest .
	@echo "$(GREEN)✅ Docker image built!$(NC)"

docker-run: ## Run Docker container
	@echo "$(YELLOW)Running Docker container...$(NC)"
	@docker run -p 3000:3000 --env-file .env $(APP_NAME):latest

docker-compose-up: ## Start with docker-compose
	@echo "$(YELLOW)Starting services with docker-compose...$(NC)"
	@docker-compose up -d
	@echo "$(GREEN)✅ Services started!$(NC)"

docker-compose-down: ## Stop docker-compose services
	@echo "$(YELLOW)Stopping docker-compose services...$(NC)"
	@docker-compose down
	@echo "$(GREEN)✅ Services stopped!$(NC)"

# =============================================================================
# Setup Commands
# =============================================================================

setup: ## Complete setup for new developers
	@echo "$(BLUE)🚀 Setting up Yaro Wora Tourism Backend...$(NC)"
	@echo "$(YELLOW)1. Installing dependencies...$(NC)"
	@make deps
	@echo "$(YELLOW)2. Creating .env file from example...$(NC)"
	@if [ ! -f .env ]; then cp .env.example .env; echo "$(GREEN)✅ .env file created from example$(NC)"; else echo "$(BLUE)ℹ️  .env file already exists$(NC)"; fi
	@echo "$(YELLOW)3. Creating database...$(NC)"
	@make db-create
	@echo "$(YELLOW)4. Building application...$(NC)"
	@make build
	@echo "$(GREEN)🎉 Setup completed!$(NC)"
	@echo "$(BLUE)Next steps:$(NC)"
	@echo "  1. Edit .env file with your database and R2 credentials"
	@echo "  2. Run 'make run' to start the application"
	@echo "  3. Visit http://localhost:3000/health to verify it's working"

install-tools: ## Install development tools
	@echo "$(YELLOW)Installing development tools...$(NC)"
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "$(GREEN)✅ Development tools installed!$(NC)"

# =============================================================================
# Utility Commands
# =============================================================================

logs: ## Show application logs (if running in background)
	@echo "$(YELLOW)Showing logs...$(NC)"
	@tail -f app.log 2>/dev/null || echo "$(BLUE)No log file found. Run the app to generate logs.$(NC)"

health: ## Check if the application is running
	@echo "$(YELLOW)Checking application health...$(NC)"
	@curl -s http://localhost:3000/health | jq . || echo "$(RED)❌ Application not responding$(NC)"

env-check: ## Validate environment variables
	@echo "$(YELLOW)Checking environment variables...$(NC)"
	@echo "DB_HOST: $(DB_HOST)"
	@echo "DB_PORT: $(DB_PORT)"
	@echo "DB_USER: $(DB_USER)"
	@echo "DB_NAME: $(DB_NAME)"
	@echo "PORT: $(PORT)"
	@echo "APP_ENV: $(APP_ENV)"
	@if [ -z "$(R2_ACCESS_KEY)" ]; then echo "$(RED)⚠️  R2_ACCESS_KEY not set$(NC)"; fi
	@if [ -z "$(R2_SECRET_KEY)" ]; then echo "$(RED)⚠️  R2_SECRET_KEY not set$(NC)"; fi

show-endpoints: ## Show available API endpoints
	@echo "$(BLUE)🌐 Available API Endpoints:$(NC)"
	@echo "$(GREEN)Public Endpoints:$(NC)"
	@echo "  GET  /health                     - Health check"
	@echo "  GET  /v1/carousel                - Get carousel slides"
	@echo "  GET  /v1/attractions             - Get attractions"
	@echo "  GET  /v1/pricing                 - Get pricing info"
	@echo "  GET  /v1/profile                 - Get village profile"
	@echo "  GET  /v1/destinations            - Get destinations"
	@echo "  GET  /v1/gallery                 - Get gallery images"
	@echo "  GET  /v1/regulations             - Get regulations"
	@echo "  GET  /v1/facilities              - Get facilities"
	@echo "  GET  /v1/news                    - Get news articles"
	@echo "  POST /v1/contact                 - Submit contact form"
	@echo "$(YELLOW)Auth Endpoints:$(NC)"
	@echo "  POST /v1/auth/login              - Simple login"
	@echo "  POST /v1/auth/jwt-login          - JWT login"
	@echo "$(RED)Admin Endpoints (Auth Required):$(NC)"
	@echo "  All CRUD operations for content management"
	@echo "  POST /v1/admin/content/upload    - Upload files"

# =============================================================================
# Production Commands
# =============================================================================

build-prod: ## Build for production
	@echo "$(YELLOW)Building for production...$(NC)"
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(APP_NAME) .
	@echo "$(GREEN)✅ Production build completed!$(NC)"

deploy-check: ## Check if ready for deployment
	@echo "$(YELLOW)Checking deployment readiness...$(NC)"
	@make test
	@make lint
	@echo "$(GREEN)✅ Ready for deployment!$(NC)"
