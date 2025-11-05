# Small Project Structure (< 10 components)

For personal projects, landing pages, simple blogs, or MVPs.

## Recommended Structure

```
project/
├── go.mod
├── main.go
├── components/          # All templ components (flat)
│   ├── layout.templ
│   ├── header.templ
│   ├── footer.templ
│   ├── home.templ
│   └── about.templ
├── handlers/            # HTTP handlers
│   ├── home.go
│   └── about.go
└── static/              # Static assets
    ├── css/
    │   └── style.css
    └── js/
        └── app.js
```

## Characteristics

- **Single `components/` directory**: All templates in one place
- **Flat structure**: No nesting, minimal folders
- **Quick navigation**: Easy to find any file
- **Simple imports**: Short import paths

## Complete Example

### 1. Project Setup

```bash
mkdir blog
cd blog

go mod init github.com/user/blog
go get github.com/a-h/templ

mkdir -p components handlers static/css static/js
```

### 2. Components

**components/layout.templ:**
```templ
package components

templ Layout(title string) {
    <!DOCTYPE html>
    <html>
        <head>
            <title>{ title }</title>
            <link rel="stylesheet" href="/static/css/style.css"/>
        </head>
        <body>
            @Header()
            <main>
                { children... }
            </main>
            @Footer()
        </body>
    </html>
}
```

**components/header.templ:**
```templ
package components

templ Header() {
    <header>
        <nav>
            <a href="/">Home</a>
            <a href="/about">About</a>
        </nav>
    </header>
}
```

**components/footer.templ:**
```templ
package components

templ Footer() {
    <footer>
        <p>&copy; 2025 My Blog</p>
    </footer>
}
```

**components/home.templ:**
```templ
package components

templ Home(posts []Post) {
    <div class="posts">
        for _, post := range posts {
            <article>
                <h2>{ post.Title }</h2>
                <p>{ post.Summary }</p>
            </article>
        }
    </div>
}
```

### 3. Handlers

**handlers/home.go:**
```go
package handlers

import (
    "net/http"
    "blog/components"
)

func Home(w http.ResponseWriter, r *http.Request) {
    posts := getPosts() // Your data logic

    components.Layout("Home")(
        components.Home(posts),
    ).Render(r.Context(), w)
}
```

**handlers/about.go:**
```go
package handlers

import (
    "net/http"
    "blog/components"
)

func About(w http.ResponseWriter, r *http.Request) {
    components.Layout("About")(
        components.About(),
    ).Render(r.Context(), w)
}
```

### 4. Main Server

**main.go:**
```go
package main

import (
    "log"
    "net/http"
    "blog/handlers"
)

func main() {
    // Static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Routes
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)

    // Start
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## File Organization Tips

### Component Naming

Use descriptive, flat names:
```
components/
├── layout.templ          ✅ Base layout
├── header.templ          ✅ Site header
├── footer.templ          ✅ Site footer
├── home-hero.templ       ✅ Homepage hero section
├── post-card.templ       ✅ Blog post card
└── contact-form.templ    ✅ Contact form
```

Avoid nesting for small projects:
```
components/
├── shared/               ❌ Unnecessary nesting
│   └── button.templ
└── pages/                ❌ Over-organization
    └── home.templ
```

### Handler Organization

Match handlers to pages:
```
handlers/
├── home.go               # GET /
├── about.go              # GET /about
├── contact.go            # GET /contact, POST /contact
└── blog.go               # GET /blog, GET /blog/:slug
```

Or group by feature:
```
handlers/
├── pages.go              # Home, About, Contact
└── blog.go               # All blog routes
```

## Import Pattern

Simple, direct imports:

```go
import (
    "project/components"
    "project/handlers"
)

// Usage
components.Layout("Title")(
    components.Header(),
    components.Home(),
).Render(ctx, w)
```

## When to Add Folders

Start adding structure when you have:
- **10+ components**: Group by type (layout/, pages/, shared/)
- **5+ handlers**: Separate by domain
- **Multiple asset types**: Organize static/ into subdirectories

## Real-World Examples

### Personal Blog

```
blog/
├── go.mod
├── main.go
├── components/
│   ├── layout.templ
│   ├── header.templ
│   ├── post-list.templ
│   ├── post-detail.templ
│   └── sidebar.templ
├── handlers/
│   ├── home.go
│   └── post.go
└── static/
    └── css/style.css
```

### Landing Page

```
landing/
├── go.mod
├── main.go
├── components/
│   ├── layout.templ
│   ├── hero.templ
│   ├── features.templ
│   ├── pricing.templ
│   └── cta.templ
├── handlers/
│   └── pages.go
└── static/
    ├── css/
    └── images/
```

### Simple Dashboard

```
dashboard/
├── go.mod
├── main.go
├── components/
│   ├── layout.templ
│   ├── sidebar.templ
│   ├── stats.templ
│   └── chart.templ
├── handlers/
│   ├── dashboard.go
│   └── api.go
└── static/
    ├── css/
    └── js/
```

## Development Workflow

### Quick Start

```bash
# Terminal 1: Watch templates
templ generate --watch

# Terminal 2: Run server
go run .
```

### With Air (Recommended)

```bash
# Install
go install github.com/cosmtrek/air@latest

# Run (auto-reload)
air
```

## Common Patterns

### Page Handler Template

```go
func PageName(w http.ResponseWriter, r *http.Request) {
    // 1. Get data
    data := getData()

    // 2. Render
    components.Layout("Title")(
        components.PageName(data),
    ).Render(r.Context(), w)
}
```

### Form Handler

```go
func ContactForm(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        components.Layout("Contact")(
            components.ContactForm(),
        ).Render(r.Context(), w)
        return
    }

    // POST: Handle form submission
    r.ParseForm()
    // ... process form
    http.Redirect(w, r, "/success", http.StatusSeeOther)
}
```

## Best Practices

1. **Keep it flat**: Don't create folders until you need them
2. **Descriptive names**: Use `post-card.templ`, not `card.templ`
3. **One handler per route**: Easy to find and modify
4. **Co-locate related code**: Handler and component for same feature nearby
5. **Use Layout wrapper**: Consistent page structure

## Anti-Patterns

❌ **Premature folders**:
```
components/
├── shared/       # Only 1-2 files
│   └── button.templ
└── layouts/      # Only 1 file
    └── base.templ
```

❌ **Unclear names**:
```
components/
├── page1.templ
├── comp.templ
└── thing.templ
```

❌ **Mixed concerns**:
```
components/
├── user.templ           # Template
├── user-service.go      # Business logic ❌
└── user-repository.go   # Data access ❌
```

## Migration Path

When growing beyond 10 components, consider:

1. **Create folders**: `components/layout/`, `components/pages/`, `components/shared/`
2. **Move files**: Group related components
3. **Update imports**: Add subfolder to import paths
4. **Use `internal/`**: Move to `internal/components/` for better organization

See [structure-medium.md](./structure-medium.md) for next steps.

## Checklist

- [ ] All components in single `components/` folder
- [ ] Handlers match routes/features
- [ ] Static files organized by type
- [ ] Descriptive file names
- [ ] Simple import paths
- [ ] No unnecessary nesting
- [ ] Ready to scale when needed
