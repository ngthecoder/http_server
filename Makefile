BINARY_NAME=http_server
MAIN_PACKAGE=./app
BUILD_DIR=./build

.PHONY: build run

build:
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(BINARY_NAME)