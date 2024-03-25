# Makefile for a go project

# Variables
GO := go
BUILD_DIR := build
BIN_NAME := tdfhf
MAIN_FILE := cmd/tdfhf/main.go

# Build target
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(BIN_NAME) $(MAIN_FILE)

# Clean target
clean:
	@rm -rf $(BUILD_DIR)

.PHONY: build clean