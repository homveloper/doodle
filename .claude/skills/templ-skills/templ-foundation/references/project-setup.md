# Templ Project Setup Guide

## Project Initialization

### Step-by-Step Setup

#### 1. Create Project Directory

```bash
mkdir my-templ-project
cd my-templ-project
```

#### 2. Initialize Go Module

```bash
go mod init github.com/username/my-templ-project
```

Replace with your actual module path.

#### 3. Install Dependencies

```bash
# Core dependency
go get github.com/a-h/templ

# Development tools
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/cosmtrek/air@latest
```

#### 4. Create Directory Structure

```bash
mkdir -p components/{layouts,pages,shared}
mkdir -p handlers
mkdir -p static/{css,js,images}
mkdir -p tmp  # For air
```

### Directory Structure

#### Minimal Structure (Small Projects)

```
project/
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── .gitignore
├── components/
│   ├── hello.templ
│   └── layout.templ
└── static/
    └── styles.css
```

#### Standard Structure (Medium Projects)

```
project/
├── cmd/
│   └── server/
│       └── main.go          # Entry point
├── internal/
│   ├── components/          # Templ components
│   │   ├── layouts/
│   │   │   ├── base.templ
│   │   │   └── nav.templ
│   │   ├── pages/
│   │   │   ├── home.templ
│   │   │   └── about.templ
│   │   └── shared/
│   │       ├── button.templ
│   │       └── card.templ
│   ├── handlers/            # HTTP handlers
│   │   ├── home.go
│   │   └── about.go
│   └── models/              # Data models
│       └── user.go
├── static/
│   ├── css/
│   │   └── main.css
│   ├── js/
│   │   └── app.js
│   └── images/
├── go.mod
├── go.sum
├── Makefile
├── .air.toml
└── .gitignore
```

#### Advanced Structure (Large Projects)

```
project/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── web/
│   │   ├── components/
│   │   │   ├── auth/        # Feature-based grouping
│   │   │   │   ├── login.templ
│   │   │   │   └── signup.templ
│   │   │   ├── dashboard/
│   │   │   │   ├── overview.templ
│   │   │   │   └── stats.templ
│   │   │   └── shared/
│   │   │       ├── button.templ
│   │   │       ├── modal.templ
│   │   │       └── table.templ
│   │   ├── layouts/
│   │   │   ├── base.templ
│   │   │   ├── auth.templ
│   │   │   └── dashboard.templ
│   │   └── pages/
│   │       ├── home.templ
│   │       ├── about.templ
│   │       └── contact.templ
│   ├── handlers/
│   │   ├── auth/
│   │   │   ├── login.go
│   │   │   └── signup.go
│   │   ├── dashboard/
│   │   │   └── dashboard.go
│   │   └── pages/
│   │       └── static.go
│   ├── services/            # Business logic
│   │   ├── auth.go
│   │   └── user.go
│   ├── models/              # Domain models
│   │   └── user.go
│   └── middleware/          # HTTP middleware
│       ├── auth.go
│       └── logging.go
├── web/
│   └── static/
│       ├── css/
│       ├── js/
│       └── images/
├── config/
│   └── config.go
├── migrations/              # Database migrations
├── go.mod
├── go.sum
├── Makefile
├── .air.toml
├── .env.example
├── docker-compose.yml
└── Dockerfile
```

## Configuration Files

### go.mod

```go
module github.com/username/project

go 1.21

require (
    github.com/a-h/templ v0.2.543
)
```

### Makefile

```makefile
.PHONY: help templ dev build test clean install

help:
	@echo "Available commands:"
	@echo "  make install  - Install dependencies"
	@echo "  make templ    - Generate templ files"
	@echo "  make dev      - Run development server"
	@echo "  make build    - Build for production"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"

install:
	@echo "Installing dependencies..."
	@go mod download
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/cosmtrek/air@latest

templ:
	@echo "Generating templ files..."
	@templ generate

dev:
	@echo "Starting development server..."
	@air

build: templ
	@echo "Building for production..."
	@go build -o dist/app ./cmd/server

test: templ
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf tmp/ dist/
	@find . -name "*_templ.go" -delete
```

### .gitignore

```gitignore
# Templ generated files
*_templ.go

# Build artifacts
tmp/
dist/
bin/
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary
*.test

# Output of the go coverage tool
*.out

# Air
.air.toml.tmp

# Dependency directories
vendor/

# Go workspace file
go.work

# Environment variables
.env
.env.local

# IDE
.vscode/
.idea/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db
```

### .air.toml

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "templ generate && go build -o ./tmp/main ./cmd/server"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go", "_templ.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "templ", "tpl", "html"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

### .env.example

```bash
# Server
PORT=8080
HOST=localhost

# Environment
ENV=development

# Database (if using)
DATABASE_URL=postgres://user:pass@localhost:5432/dbname

# Session secret (if using sessions)
SESSION_SECRET=your-secret-key-here
```

## IDE Configuration

### VS Code

#### Extensions

Install "templ" extension:
1. Open VS Code
2. Go to Extensions (Cmd+Shift+X)
3. Search for "templ"
4. Install "templ" by a-h

#### settings.json

```json
{
  "files.associations": {
    "*.templ": "templ"
  },
  "templ.lsp.enabled": true,
  "[templ]": {
    "editor.defaultFormatter": "a-h.templ"
  }
}
```

#### tasks.json

```json
{
  "version": "2.0.0",
  "tasks": [
    {
      "label": "templ generate",
      "type": "shell",
      "command": "templ generate",
      "group": {
        "kind": "build",
        "isDefault": true
      }
    },
    {
      "label": "run dev server",
      "type": "shell",
      "command": "make dev",
      "group": {
        "kind": "test",
        "isDefault": true
      }
    }
  ]
}
```

### GoLand / IntelliJ IDEA

1. Install "Templ" plugin from marketplace
2. Enable Go Template support
3. Configure file associations:
   - Settings → Editor → File Types
   - Add `*.templ` pattern to "Go Template files"

### Neovim

#### Install treesitter-templ

```lua
-- In your nvim config
require'nvim-treesitter.configs'.setup {
  ensure_installed = {
    "templ",
    "go",
    "html",
    -- other languages
  },
}
```

#### LSP Configuration

```lua
-- Setup templ LSP
require'lspconfig'.templ.setup{}
```

## Docker Setup (Optional)

### Dockerfile

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Generate templ files
RUN templ generate

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

# Run stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary and static files
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

EXPOSE 8080

CMD ["./server"]
```

### docker-compose.yml

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - ENV=production
      - PORT=8080
    volumes:
      - ./static:/root/static:ro
    restart: unless-stopped
```

## Environment Management

### config/config.go

```go
package config

import (
    "os"
    "strconv"
)

type Config struct {
    Port        int
    Host        string
    Environment string
}

func Load() *Config {
    return &Config{
        Port:        getEnvAsInt("PORT", 8080),
        Host:        getEnv("HOST", "localhost"),
        Environment: getEnv("ENV", "development"),
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}

func getEnvAsInt(key string, fallback int) int {
    if value := os.Getenv(key); value != "" {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return fallback
}
```

## Quick Start Commands

```bash
# Clone or create project
mkdir my-project && cd my-project

# Initialize
make install

# Start development
make dev

# Build for production
make build

# Run production build
./dist/app
```

## Common Issues

### Issue: "templ: command not found"

**Solution**: Add GOPATH/bin to PATH
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Issue: Generated files not found

**Solution**: Run templ generate before building
```bash
templ generate && go run .
```

### Issue: Import path errors

**Solution**: Ensure go.mod module path matches imports
```go
// go.mod
module github.com/user/project

// Import as
import "github.com/user/project/components"
```

### Issue: Air not reloading

**Solution**: Check .air.toml includes .templ files
```toml
include_ext = ["go", "templ"]
```
