# Define variables
APP_NAME = loan-service
GO_FILES = $(shell find . -name "*.go")
BIN_DIR = bin
BIN_NAME = $(BIN_DIR)/$(APP_NAME)

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build: $(GO_FILES)
	@echo "Building the application..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_NAME) cmd/loan-service/main.go

# Run the application
.PHONY: run
run: build
	@echo "Running the application..."
	@./$(BIN_NAME)

# Install dependencies
.PHONY: install
install:
	@echo "Installing dependencies..."
	@go mod tidy

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	@rm -rf $(BIN_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Run linting (optional, using golangci-lint for example)
.PHONY: lint
lint:
	@echo "Running linting..."
	@golangci-lint run

# Run the server with environment variables loaded from .env (optional)
.PHONY: run-env
run-env: build
	@echo "Running the server with environment variables from .env..."
	@source .env && ./$(BIN_NAME)
