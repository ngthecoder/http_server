BINARY_NAME=http_server

MAIN_PACKAGE=./app

BUILD_DIR=./build

build:
	@echo "Building the binary..."
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(BINARY_NAME)