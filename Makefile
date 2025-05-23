# Project settings
BINARY_NAME = zhconvert
CMD_DIR = ./cmd/zhconvert
BUILD_DIR = ./build

.PHONY: all build clean install run help

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	chmod +x $(BUILD_DIR)/$(BINARY_NAME)

install:
	go install $(CMD_DIR)

run:
	go run $(CMD_DIR)/main.go

clean:
	rm -f $(BUILD_DIR)/*

help:
	@echo "Makefile commands:"
	@echo "  make           - Build the CLI binary"
	@echo "  make build     - Build the binary to $(BUILD_DIR)/$(BINARY_NAME)"
	@echo "  make install   - Install to \$GOBIN"
	@echo "  make run       - Run directly with go run"
	@echo "  make clean     - Remove binary"
