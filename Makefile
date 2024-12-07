# List of packages to test
PACKAGES := set list

# Base path for the go-collection directory
BASE_PATH := $(shell pwd)

# Output directory for build artifacts
BUILD_DIR := $(BASE_PATH)/build

# Ensure build directory exists
$(shell mkdir -p $(BUILD_DIR))

.PHONY: test
test:
	# Loop through each package and run tests with coverage
	@for pkg in $(PACKAGES); do \
		go test -coverprofile=$(BUILD_DIR)/coverage_$$pkg.out $(BASE_PATH)/$$pkg; \
		go tool cover -html=$(BUILD_DIR)/coverage_$$pkg.out -o $(BUILD_DIR)/coverage_$$pkg.html; \
		echo "Coverage report for $$pkg generated at $(BUILD_DIR)/coverage_$$pkg.html"; \
	done

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_DIR)/coverage_*.out $(BUILD_DIR)/coverage_*.html
