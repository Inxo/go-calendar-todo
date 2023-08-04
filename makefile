# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the binary output file
BINARY_NAME=calendar_app
BUILD_DIR=build

# Default target: build the application
all: build

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .

# Clean up the generated binary and build directory
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Install necessary dependencies
deps:
	$(GOGET) -u github.com/mattn/go-sqlite3
	$(GOGET) -u github.com/gdamore/tcell

# Run the application
run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .
	./$(BUILD_DIR)/$(BINARY_NAME)

# Install and run the application
install: deps
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) .
	./$(BUILD_DIR)/$(BINARY_NAME)

# Test the application
test:
	$(GOTEST) ./...

# Format the source code
fmt:
	$(GOCMD) fmt ./...

# Run lint checks
lint:
	golangci-lint run

# Help command to display available targets
help:
	@echo "Available targets:"
	@echo "  build    : Build the application (outputs to build directory)"
	@echo "  clean    : Clean up the generated binary and build directory"
	@echo "  deps     : Install necessary dependencies"
	@echo "  run      : Build and run the application"
	@echo "  install  : Install dependencies and run the application"
	@echo "  test     : Run tests"
	@echo "  fmt      : Format the source code"
	@echo "  lint     : Run lint checks"
