.PHONY: help install templ dev build test clean

help:
	@echo "Available commands:"
	@echo "  make install  - Install dependencies and tools"
	@echo "  make templ    - Generate Go files from .templ"
	@echo "  make dev      - Run development server with live reload"
	@echo "  make build    - Build for production"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"

install:
	@echo "📦 Installing dependencies..."
	@go mod download
	@echo "🔧 Installing templ CLI..."
	@go install github.com/a-h/templ/cmd/templ@latest
	@echo "🌪️  Installing air (live reload)..."
	@go install github.com/cosmtrek/air@latest
	@echo "✅ Installation complete!"

templ:
	@echo "🔄 Generating templ files..."
	@templ generate

dev:
	@echo "🚀 Starting development server..."
	@air

build: templ
	@echo "🏗️  Building for production..."
	@go build -o dist/app .
	@echo "✅ Build complete: dist/app"

test: templ
	@echo "🧪 Running tests..."
	@go test -v ./...

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf tmp/ dist/
	@find . -name "*_templ.go" -delete
	@echo "✅ Clean complete!"
