.PHONY: help dev dev-clean dev-info build test clean run install-deps templ-generate templ-watch check-deps kill-port logs watch

# Default target
.DEFAULT_GOAL := help

# Variables
APP_NAME := shop-server
PORT := 8080

## help: Show this help message
help:
	@echo "Available commands:"
	@echo ""
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/^## /  /'
	@echo ""

## dev: Start development server with hot reload (Air)
dev: check-deps
	@echo "🚀 Starting development server with hot reload..."
	@echo "📱 Open http://localhost:$(PORT) in your browser"
	@echo "💡 Changes to .go and .templ files will trigger automatic rebuild"
	@echo "⏱️  Initial build may take a few seconds..."
	@echo ""
	@air

## dev-clean: Clean and start development server
dev-clean: clean kill-port dev

## kill-port: Kill any process running on PORT
kill-port:
	@echo "🔪 Checking port $(PORT)..."
	@lsof -ti:$(PORT) | xargs kill -9 2>/dev/null || echo "✅ Port $(PORT) is free"

## check-deps: Check if required tools are installed
check-deps:
	@echo "🔍 Checking dependencies..."
	@command -v templ >/dev/null 2>&1 || (echo "❌ templ not found. Run: make install-deps" && exit 1)
	@command -v air >/dev/null 2>&1 || (echo "❌ air not found. Run: make install-deps" && exit 1)
	@echo "✅ All dependencies installed"

## build: Build the application
build: templ-generate
	@echo "🔨 Building application..."
	@go build -o $(APP_NAME) .
	@echo "✅ Build complete: ./$(APP_NAME)"

## run: Build and run the application
run: build
	@echo "🚀 Running application..."
	@./$(APP_NAME)

## test: Run all tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...

## test-cover: Run tests with coverage
test-cover:
	@echo "🧪 Running tests with coverage..."
	@go test -cover ./...
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

## clean: Clean build artifacts and temporary files
clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(APP_NAME)
	@rm -rf tmp/
	@rm -f build-errors.log
	@rm -f coverage.out coverage.html
	@find . -name "*_templ.go" -type f -delete
	@echo "✅ Clean complete"

## install-deps: Install development dependencies
install-deps:
	@echo "📦 Installing dependencies..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/air-verse/air@latest
	@go mod download
	@echo "✅ Dependencies installed"

## templ-generate: Generate Go code from Templ templates
templ-generate:
	@echo "⚡ Generating Templ templates..."
	@templ generate
	@echo "✅ Templ generation complete"

## templ-watch: Watch and auto-generate Templ templates
templ-watch:
	@echo "👀 Watching Templ templates for changes..."
	@templ generate --watch

## fmt: Format code
fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...
	@templ fmt .
	@echo "✅ Code formatted"

## lint: Run linters (requires golangci-lint)
lint:
	@echo "🔍 Running linters..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not installed. Install with:"; \
		echo "    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

## mod-tidy: Tidy go.mod and go.sum
mod-tidy:
	@echo "📦 Tidying modules..."
	@go mod tidy
	@echo "✅ Modules tidied"

## init: Initialize project (install deps + generate templates)
init: install-deps templ-generate
	@echo "✅ Project initialized successfully!"
	@echo ""
	@echo "Quick start:"
	@echo "  make dev    - Start development server"
	@echo "  make test   - Run tests"
	@echo "  make build  - Build for production"

## logs: Show Air build logs
logs:
	@if [ -f build-errors.log ]; then \
		echo "📋 Build error logs:"; \
		cat build-errors.log; \
	else \
		echo "✅ No error logs found"; \
	fi

## watch: Watch mode info
watch:
	@echo "👀 Air is watching these file types:"
	@echo "  - .go files (Go source)"
	@echo "  - .templ files (Templates)"
	@echo ""
	@echo "📁 Watching directories:"
	@echo "  - models/"
	@echo "  - handlers/"
	@echo "  - templates/"
	@echo "  - . (root)"
	@echo ""
	@echo "🔄 On change, Air will:"
	@echo "  1. Run 'templ generate'"
	@echo "  2. Build the application"
	@echo "  3. Restart the server"
	@echo ""
	@echo "Start watching with: make dev"

## dev-info: Show development server information
dev-info:
	@echo "📊 Development Server Info"
	@echo "=========================="
	@echo "App Name:    $(APP_NAME)"
	@echo "Port:        $(PORT)"
	@echo "URL:         http://localhost:$(PORT)"
	@echo "Hot Reload:  Enabled (Air)"
	@echo ""
	@echo "📁 Project Structure:"
	@ls -lh | grep '^d' || true
	@echo ""
	@echo "📦 Go Modules:"
	@go list -m all 2>/dev/null | head -5 || echo "Run 'go mod download' first"
