.PHONY: build run lint

BINARY_NAME=image-gen-mcp
BUILD_DIR=bin

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/server

run: build
	$(BUILD_DIR)/$(BINARY_NAME)

lint:
	go tool golangci-lint run --fix