.PHONY: help proto build test run-dev run-docker clean

help:
	@echo "Available commands:"
	@echo "  proto       - Generate protobuf files"
	@echo "  build       - Build all services"
	@echo "  test        - Run tests"
	@echo "  run-dev     - Run services in development mode"
	@echo "  run-docker  - Run services using Docker Compose"
	@echo "  clean       - Clean build artifacts"

proto:
	@echo "🔧 Generating protobuf files..."
	@chmod +x scripts/generate-proto.sh
	@./scripts/generate-proto.sh

build: proto
	@echo "🏗️  Building services..."
	@go build -o bin/auth-service ./services/auth-service
	@go build -o bin/user-service ./services/user-service
	@go build -o bin/asset-service ./services/asset-service
	@go build -o bin/ticket-service ./services/ticket-service
	@go build -o bin/assignment-service ./services/assignment-service
	@go build -o bin/notification-service ./services/notification-service
	@go build -o bin/bff ./services/bff

test:
	@echo "🧪 Running tests..."
	@go test -v ./...

run-dev:
	@echo "🚀 Starting services in development mode..."
	@make build
	@echo "Starting auth service..."
	@./bin/auth-service &
	@echo "Starting user service..."
	@./bin/user-service &
	@echo "Starting asset service..."
	@./bin/asset-service &
	@echo "Starting ticket service..."
	@./bin/ticket-service &
	@echo "Starting assignment service..."
	@./bin/assignment-service &
	@echo "Starting notification service..."
	@./bin/notification-service &
	@echo "Starting BFF..."
	@./bin/bff &
	@echo "All services started!"

run-docker:
	@echo "🐳 Starting services with Docker Compose..."
	@docker-compose up --build

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f proto/**/*.pb.go
	@docker-compose down
	@docker system prune -f

install-deps:
	@echo "📦 Installing dependencies..."
	@go mod tidy
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

dev-setup: install-deps proto
	@echo "🛠️  Development setup complete!"
	@echo "You can now run 'make run-dev' to start all services" 