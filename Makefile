# Ephemeral Port Exporter

APP_NAME = ephemeral-port-exporter
BUILD_DIR = bin

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod
GOFMT = gofmt
GOVET = $(GOCMD) vet

# Binary name
BINARY_NAME = $(APP_NAME)
BINARY_UNIX = $(BINARY_NAME)_unix

# Default target
.PHONY: all
all: clean build

# Dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	$(GOGET) -d -v .
	$(GOMOD) tidy
	@echo "Dependencies updated"

# Build
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v ./cmd/ephemeral-port-exporter
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

build-linux:
	@echo "Building $(APP_NAME) for Linux..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) -v ./cmd/ephemeral-port-exporter
	@echo "Linux build complete: $(BUILD_DIR)/$(BINARY_UNIX)"

# Install
.PHONY: install
install:
	@echo "Installing $(APP_NAME)..."
	$(GOCMD) install .
	@echo "Installation complete"

# Run
.PHONY: run
run: build
	@echo "Running built binary..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Test
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Clean
.PHONY: clean
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete"

# Docker
# Create docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .
	@echo "Docker image built: $(APP_NAME):latest"

# Run in docker
.PHONY: docker-run
docker-run:
	@echo "Running in Docker..."
	docker run -p 2112:2112 --name $(APP_NAME) --rm $(APP_NAME):latest

# Run with Docker Compose
.PHONY: compose-up
compose-up:
	@echo "Docker Compose Up..."
	docker-compose up --build -d

.PHONY: compose-down
compose-down:
	@echo "Docker Compose Down..."
	docker-compose down

.PHONY: compose-rebuild
compose-rebuild:
	@echo "Docker Compose Rebuild..."
	docker-compose down
	docker-compose up --build -d