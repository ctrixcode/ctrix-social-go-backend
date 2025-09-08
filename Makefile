# Define the application name and build path
APP_NAME=ctrix-social-go-backend
CMD_DIR=./cmd/api

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	go build -o $(APP_NAME) $(CMD_DIR)

# Run the application
.PHONY: run
run: build
	./$(APP_NAME)

# Clean up build artifacts
.PHONY: clean
clean:
	go clean -cache
	rm -f $(APP_NAME)

# Run tests
.PHONY: test
test:
	go test ./...

# Force run a target even if a file with the same name exists
.PHONY: all build run clean test