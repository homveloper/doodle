# Makefile for Templ Projects

Automate common tasks with Make.

## Basic Makefile

Simple automation for small to medium projects:

```makefile
.PHONY: templ dev build test clean

# Generate templ files
templ:
	@echo "Generating templ files..."
	@templ generate

# Development server (requires air)
dev:
	@echo "Starting development server..."
	@air

# Build for production
build: templ
	@echo "Building..."
	@go build -o dist/app .

# Run tests
test: templ
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf dist/
	@find . -name "*_templ.go" -delete
```

**Usage:**
```bash
make templ    # Generate templates
make dev      # Start dev server with live reload
make build    # Build production binary
make test     # Run tests
make clean    # Remove generated files
```

## Advanced Makefile

More targets for larger projects:

```makefile
.PHONY: all install templ watch dev build test coverage clean help

# Default target
all: build

# Install dependencies and tools
install:
	@echo "Installing dependencies..."
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest

# Generate templ files
templ:
	@templ generate

# Watch templ files (manual mode)
watch:
	@templ generate --watch

# Development with air (auto-reload)
dev:
	@air

# Build for production
build: templ
	@echo "Building production binary..."
	@go build -ldflags="-s -w" -o dist/app .

# Run all tests
test: templ
	@go test -v ./...

# Run tests with coverage
coverage: templ
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Clean all build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf dist/ coverage.out coverage.html
	@find . -name "*_templ.go" -delete

# Show help
help:
	@echo "Available targets:"
	@echo "  make install   - Install dependencies and tools"
	@echo "  make templ     - Generate templ files"
	@echo "  make watch     - Watch and regenerate templ files"
	@echo "  make dev       - Start development server with air"
	@echo "  make build     - Build production binary"
	@echo "  make test      - Run tests"
	@echo "  make coverage  - Run tests with coverage report"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make help      - Show this help message"
```

**Usage:**
```bash
make install     # One-time setup
make dev         # Daily development
make test        # Before committing
make coverage    # Check test coverage
make build       # Production build
make clean       # Clean up
```

## Common Patterns

### Multi-platform Build

```makefile
.PHONY: build-all

build-all: templ
	@echo "Building for multiple platforms..."
	@GOOS=linux GOARCH=amd64 go build -o dist/app-linux-amd64 .
	@GOOS=darwin GOARCH=amd64 go build -o dist/app-darwin-amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o dist/app-darwin-arm64 .
	@GOOS=windows GOARCH=amd64 go build -o dist/app-windows-amd64.exe .
	@echo "Built for Linux, macOS (Intel/ARM), and Windows"
```

### Conditional Air Check

```makefile
dev:
	@which air > /dev/null || (echo "Air not installed. Run: make install" && exit 1)
	@air
```

### Version from Git

```makefile
VERSION := $(shell git describe --tags --always --dirty)

build: templ
	@go build -ldflags="-s -w -X main.version=$(VERSION)" -o dist/app .
```

### Parallel Test Execution

```makefile
test-parallel: templ
	@go test -v -parallel=4 ./...
```

## Tips

### 1. Always Generate Before Build

Use `build: templ` dependency to ensure templates are generated:

```makefile
build: templ
	@go build -o dist/app .
```

### 2. Silent Commands

Use `@` prefix to hide command echo:

```makefile
templ:
	@templ generate        # Silent
	templ generate         # Shows: templ generate
```

### 3. Error Handling

Stop on error with proper exit codes:

```makefile
test: templ
	@go test ./... || exit 1
```

### 4. Variables for Reusability

```makefile
APP_NAME := myapp
BUILD_DIR := dist
BINARY := $(BUILD_DIR)/$(APP_NAME)

build: templ
	@go build -o $(BINARY) .
```

### 5. Check Tool Installation

```makefile
check-templ:
	@which templ > /dev/null || (echo "Error: templ not installed" && exit 1)

build: check-templ templ
	@go build -o dist/app .
```

## Integration with Go Commands

### Run with Arguments

```makefile
run: templ
	@go run . $(ARGS)

# Usage: make run ARGS="--port 8080"
```

### Build Tags

```makefile
build-dev: templ
	@go build -tags=dev -o dist/app .

build-prod: templ
	@go build -tags=prod -o dist/app .
```

### Verbose Output

```makefile
test-verbose: templ
	@go test -v -count=1 ./...
```

## Troubleshooting

### Make Not Found

Install Make:
- **macOS**: `xcode-select --install`
- **Linux**: Usually pre-installed, or `apt install make`
- **Windows**: Use WSL or `choco install make`

### Target Not Running

Check `.PHONY` declaration:

```makefile
.PHONY: build test clean    # Declare non-file targets
```

### Commands Not Working

Makefile requires **tabs**, not spaces:

```makefile
build:
→   @go build .    # ← This must be a tab character
```

### Parallel Execution Issues

Use `.NOTPARALLEL` if needed:

```makefile
.NOTPARALLEL:    # Force sequential execution
```

## Example: Complete Project Makefile

```makefile
.PHONY: all install templ dev build test clean

APP := myapp
BUILD_DIR := dist
VERSION := $(shell git describe --tags --always --dirty)

all: build

install:
	@echo "Installing dependencies and tools..."
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest

templ:
	@templ generate

dev:
	@which air > /dev/null || (echo "Run 'make install' first" && exit 1)
	@air

build: templ
	@echo "Building $(APP) version $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(BUILD_DIR)/$(APP) .
	@echo "Binary: $(BUILD_DIR)/$(APP)"

test: templ
	@go test -v -race ./...

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@find . -name "*_templ.go" -delete

help:
	@echo "Targets:"
	@echo "  install - Install dependencies"
	@echo "  dev     - Start development server"
	@echo "  build   - Build production binary"
	@echo "  test    - Run tests"
	@echo "  clean   - Remove artifacts"
```

## Best Practices

1. **Use .PHONY**: Declare all non-file targets
2. **Dependencies**: Use target dependencies (`build: templ`)
3. **Silent by default**: Use `@` for cleaner output
4. **Help target**: Always include a help target
5. **Error handling**: Exit with non-zero on errors
6. **Variables**: Use variables for paths and names
7. **Check tools**: Verify required tools are installed

## Alternatives to Make

If Make is not available or preferred:

### Bash Script (run.sh)

```bash
#!/bin/bash

case "$1" in
  "templ")
    templ generate
    ;;
  "dev")
    air
    ;;
  "build")
    templ generate && go build -o dist/app .
    ;;
  *)
    echo "Usage: ./run.sh {templ|dev|build}"
    exit 1
    ;;
esac
```

### Go Task (Taskfile.yml)

```yaml
version: '3'

tasks:
  templ:
    cmds:
      - templ generate

  dev:
    cmds:
      - air

  build:
    deps: [templ]
    cmds:
      - go build -o dist/app .
```

### Just (justfile)

```just
templ:
    templ generate

dev:
    air

build: templ
    go build -o dist/app .
```

But **Makefile is most common** and widely supported.
