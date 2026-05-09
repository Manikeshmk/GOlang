.PHONY: help build run test clean docker-build docker-up docker-down lint fmt

# Variables
BINARY_NAME=summarizer-api
DOCKER_IMAGE=silent-meeting-summarizer
DOCKER_TAG=latest
PORT=8080

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) cmd/api/main.go

run: ## Run the application
	@echo "Running $(BINARY_NAME)..."
	@if [ ! -f .env ]; then cp .env.example .env; fi
	go run cmd/api/main.go

dev: ## Run with hot reload (requires air)
	@which air > /dev/null || go install github.com/cosmtrek/air@latest
	air

test: ## Run tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-unit: ## Run unit tests
	go test -v -short ./tests/unit/...

test-integration: ## Run integration tests
	docker-compose -f deployments/docker/docker-compose.test.yml up -d
	go test -v ./tests/integration/...
	docker-compose -f deployments/docker/docker-compose.test.yml down

coverage: ## Generate coverage report
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f bin/$(BINARY_NAME)
	rm -f coverage.out coverage.html

lint: ## Run golangci-lint
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

fmt: ## Format code
	go fmt ./...
	goimports -w .

vet: ## Run go vet
	go vet ./...

docker-build: ## Build Docker image
	@echo "Building Docker image $(DOCKER_IMAGE):$(DOCKER_TAG)..."
	docker build -f deployments/docker/Dockerfile -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-up: ## Start Docker Compose
	@echo "Starting Docker containers..."
	docker-compose -f deployments/docker/docker-compose.yml up -d

docker-down: ## Stop Docker Compose
	@echo "Stopping Docker containers..."
	docker-compose -f deployments/docker/docker-compose.yml down

docker-logs: ## View Docker logs
	docker-compose -f deployments/docker/docker-compose.yml logs -f api

docker-clean: ## Clean Docker images and volumes
	docker-compose -f deployments/docker/docker-compose.yml down -v
	docker rmi $(DOCKER_IMAGE):$(DOCKER_TAG) || true

docker-shell: ## Open shell in API container
	docker-compose -f deployments/docker/docker-compose.yml exec api sh

install-deps: ## Install dependencies
	go mod download
	go mod tidy

security-check: ## Run security scan with gosec
	@which gosec > /dev/null || go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec ./...

migrate-up: ## Run database migrations
	@echo "Running migrations..."
	go run cmd/api/main.go

docs: ## Generate API documentation
	@which swag > /dev/null || go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g cmd/api/main.go

version: ## Show version info
	@echo "Go version: $$(go version)"
	@echo "Binary: $(BINARY_NAME)"

all: clean lint test build ## Run all checks and build

.DEFAULT_GOAL := help
