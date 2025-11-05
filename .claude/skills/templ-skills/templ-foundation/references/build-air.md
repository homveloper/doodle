# Air - Live Reload for Go

Automatic server restart on file changes during development.

## What is Air?

[Air](https://github.com/cosmtrek/air) is a live reload tool for Go applications. When you save a `.go` or `.templ` file, Air automatically:
1. Runs `templ generate` (if configured)
2. Rebuilds your Go binary
3. Restarts the server

**Benefits:**
- No manual server restarts
- Faster development feedback loop
- Watches multiple file types

## Installation

```bash
go install github.com/cosmtrek/air@latest
```

Verify installation:
```bash
air -v
```

## Basic Configuration

Create `.air.toml` in project root:

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
  exclude_dir = ["tmp", "vendor"]
  delay = 1000
```

**Run:**
```bash
air
```

Air will:
- Watch `.go` and `.templ` files
- Run `templ generate && go build` on changes
- Restart server automatically

## Minimal Configuration

Bare minimum `.air.toml`:

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
```

This is enough for most projects.

## Complete Configuration

Full `.air.toml` with all options:

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  # Command to build
  cmd = "templ generate && go build -o ./tmp/main ."

  # Binary path
  bin = "tmp/main"

  # Watch these file extensions
  include_ext = ["go", "templ", "html", "css", "js"]

  # Exclude these directories
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]

  # Exclude files matching these patterns
  exclude_regex = ["_test.go"]

  # Delay before rebuild (ms)
  delay = 1000

  # Stop on build error
  stop_on_error = true

  # Log file for build errors
  log = "build-errors.log"

  # Kill delay (time to wait before killing old process)
  kill_delay = "500ms"

[color]
  main = "magenta"
  watcher = "cyan"
  build = "yellow"
  runner = "green"

[log]
  time = false

[misc]
  clean_on_exit = true
```

## Common Use Cases

### Templ + Static Files

Watch templates and static assets:

```toml
[build]
  cmd = "templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ", "css", "js"]
  exclude_dir = ["tmp", "vendor", "node_modules"]
```

### With Build Tags

Development vs production tags:

```toml
[build]
  cmd = "templ generate && go build -tags=dev -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
```

### Multiple Commands

Run linter before build:

```toml
[build]
  cmd = "golangci-lint run && templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
```

### Custom Binary Name

```toml
[build]
  cmd = "templ generate && go build -o ./tmp/myapp ."
  bin = "tmp/myapp"
  include_ext = ["go", "templ"]
```

## Usage

### Start Air

```bash
air
```

### With Custom Config

```bash
air -c .air.custom.toml
```

### Debug Mode

```bash
air -d
```

Shows detailed logs of what Air is doing.

## Integration with Makefile

```makefile
.PHONY: dev

dev:
	@which air > /dev/null || (echo "Air not installed. Run: go install github.com/cosmtrek/air@latest" && exit 1)
	@air
```

Usage:
```bash
make dev
```

## Gitignore

Add to `.gitignore`:

```gitignore
# Air
tmp/
.air.toml.tmp
build-errors.log
```

## Troubleshooting

### Air Not Detecting Changes

**Problem:** Changes to files not triggering rebuild.

**Solutions:**

1. **Check include_ext:**
   ```toml
   include_ext = ["go", "templ"]  # Make sure "templ" is included
   ```

2. **Check exclude_dir:**
   ```toml
   exclude_dir = ["tmp", "vendor"]  # Don't exclude your source dirs
   ```

3. **Try poll mode** (for network drives):
   ```toml
   [build]
     poll = true
     poll_interval = 500  # ms
   ```

### Build Errors Not Showing

**Problem:** Build fails silently.

**Solution:** Enable build error log:

```toml
[build]
  log = "build-errors.log"
  stop_on_error = true
```

Check `build-errors.log` for details.

### Process Not Killed

**Problem:** Old process keeps running.

**Solution:** Increase kill delay:

```toml
[build]
  kill_delay = "1s"  # Increase from default
```

### Too Many Rebuilds

**Problem:** Air rebuilds too frequently.

**Solution:** Increase delay:

```toml
[build]
  delay = 2000  # Wait 2 seconds before rebuild
```

### Templ Generate Not Running

**Problem:** Generated files not updating.

**Solution:** Ensure `templ generate` is in cmd:

```toml
[build]
  cmd = "templ generate && go build -o ./tmp/main ."
```

Verify templ is in PATH:
```bash
which templ
```

## Tips

### 1. Exclude Test Files

Avoid rebuilding on test file changes:

```toml
[build]
  exclude_regex = ["_test.go"]
```

### 2. Build Tags for Development

```toml
[build]
  cmd = "templ generate && go build -tags=dev -o ./tmp/main ."
```

Then in code:
```go
//go:build dev
// +build dev

// Development-only code
```

### 3. Clear Console on Rebuild

```toml
[screen]
  clear_on_rebuild = true
```

### 4. Colored Output

```toml
[color]
  main = "magenta"
  watcher = "cyan"
  build = "yellow"
  runner = "green"
```

### 5. Multiple Projects

Use different configs:

```bash
air -c .air.api.toml      # API server
air -c .air.worker.toml   # Background worker
```

## Alternative: Manual Watch

If Air is too heavy, use manual watch:

**Terminal 1** - Watch templates:
```bash
templ generate --watch
```

**Terminal 2** - Run server:
```bash
go run .
```

Manually restart Terminal 2 when Go code changes.

**Pros:** Simple, no extra tools
**Cons:** Manual server restart needed

## Comparison: Air vs Manual

| Feature | Air | Manual (`templ --watch`) |
|---------|-----|--------------------------|
| Auto-restart | ✅ Yes | ❌ Manual |
| Watches .go | ✅ Yes | ❌ No |
| Watches .templ | ✅ Yes | ✅ Yes |
| Setup | .air.toml | None |
| Speed | Fast | Fast |
| Resource usage | Low | Lower |

**Recommendation:** Use Air for daily development.

## Example: Complete Setup

**1. Install Air:**
```bash
go install github.com/cosmtrek/air@latest
```

**2. Create `.air.toml`:**
```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "templ generate && go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "templ"]
  exclude_dir = ["tmp", "vendor"]
  delay = 1000
  stop_on_error = true

[log]
  time = false

[misc]
  clean_on_exit = true
```

**3. Add to `.gitignore`:**
```gitignore
tmp/
.air.toml.tmp
```

**4. Run:**
```bash
air
```

**5. Develop:**
- Edit `.templ` or `.go` files
- Save
- Air automatically rebuilds and restarts
- Refresh browser

## Best Practices

1. **Use Air in development**: Fast feedback loop
2. **Exclude test files**: Avoid unnecessary rebuilds
3. **Set appropriate delay**: Balance speed vs stability
4. **Clean on exit**: Remove tmp files automatically
5. **Version control**: Commit `.air.toml`, ignore `tmp/`
6. **Stop on error**: Catch build issues immediately

## When Not to Use Air

- **Production**: Never use Air in production
- **CI/CD**: Build normally in pipelines
- **Testing**: Use `go test` directly
- **Debugging**: Use IDE debugger or `dlv`

Air is **development-only** tool for fast iteration.
