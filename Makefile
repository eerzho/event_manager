include .env
export

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

deps: ## Update and download module dependencies
	@echo "Updating module dependencies"
	go mod tidy
	@echo "Downloading module dependencies"
	go mod download
.PHONY: deps

swag-v1: ## Generate Swagger documentation for v1
	@echo "Checking for swag"
	@if ! command -v swag &> /dev/null; then \
		echo "swag not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	else \
		echo "swag is already installed."; \
	fi
	@echo "Generating Swagger documentation"
	swag init -g internal/handler/http/v1/v1.go
.PHONY: swag-v1

run: deps swag-v1 ## Run the application
	@echo "Running the application"
	go run ./cmd/app
.PHONY: run
