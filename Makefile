# Blog Aggregator Makefile

# Database configuration
DB_URL ?= postgres://postgres:root@localhost:5433/gator?sslmode=disable
MIGRATION_DIR = sql/schema

# Build
build:
	go build -o bin/gator .

# Run the application
run:
	go run .

# Install dependencies
deps:
	go mod tidy
	go mod download

# Database migrations
migrate-up:
	goose postgres "$(DB_URL)" -dir $(MIGRATION_DIR) up

migrate-down:
	goose postgres "$(DB_URL)" -dir $(MIGRATION_DIR) down

migrate-status:
	goose postgres "$(DB_URL)" -dir $(MIGRATION_DIR) status

migrate-reset:
	goose postgres "$(DB_URL)" -dir $(MIGRATION_DIR) reset

# Generate sqlc code
sqlc:
	sqlc generate

# Clean build artifacts
clean:
	rm -rf bin/

# Test
test:
	go test ./...

# Development helpers
dev: deps sqlc build

.PHONY: build run deps migrate-up migrate-down migrate-status migrate-reset sqlc clean test dev
