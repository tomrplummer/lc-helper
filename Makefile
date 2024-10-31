BINARY_NAME=lc-helper
OUTPUT_DIR=./bin

# Ensure the output directory exists
$(OUTPUT_DIR):
	mkdir -p $(OUTPUT_DIR)

# Build for each platform
build-windows: $(OUTPUT_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows-amd64.exe ./cli

build-linux: $(OUTPUT_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux-amd64 ./cli

build-mac-amd64: $(OUTPUT_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-amd64 ./cli

build-mac-arm64: $(OUTPUT_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-darwin-arm64 ./cli

# Build all platforms
build-all: build-windows build-linux build-mac-amd64 build-mac-arm64

# Clean up binaries
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: build-windows build-linux build-mac-amd64 build-mac-arm64 build-all clean
