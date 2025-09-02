# Makefile for the flypro-assessment project

# Go parameters
BINARY_NAME=flypro-assessment
CMD_PATH=./cmd/server

# Goose parameters
MIGRATIONS_DIR=./migrations

.PHONY: all build run test clean

all: build

# Build the Go application
build:
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) $(CMD_PATH)

# Run the Go application
run: build
	@echo "Starting the application..."
	@./$(BINARY_NAME)

# Run the tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean the binary
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Run goimports to format imports
imports:
	@echo "Running goimports..."
	@goimports -local flypro-assessment -w .

# Goose migrations
.PHONY: migrate-up migrate-down migrate-status migrate-create

migrate-up:
	@echo "Running migrations up..."
	@goose -dir $(MIGRATIONS_DIR) up

migrate-down:
	@echo "Running migrations down..."
	@goose -dir $(MIGRATIONS_DIR) down

migrate-status:
	@echo "Checking migration status..."
	@goose -dir $(MIGRATIONS_DIR) status

migrate-create:
	@echo "Creating new migration..."
	@goose -dir $(MIGRATIONS_DIR) create $(name) sql
