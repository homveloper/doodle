# Medium Project Structure (10-50 components)

For SaaS products, e-commerce sites, content management systems, or growing applications.

## Recommended Structure

```
project/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── components/
│   │   ├── layout/
│   │   │   ├── base.templ
│   │   │   ├── header.templ
│   │   │   └── footer.templ
│   │   ├── pages/
│   │   │   ├── home.templ
│   │   │   ├── about.templ
│   │   │   ├── contact.templ
│   │   │   └── dashboard.templ
│   │   └── shared/
│   │       ├── button.templ
│   │       ├── card.templ
│   │       ├── form.templ
│   │       └── modal.templ
│   ├── handlers/
│   │   ├── pages.go
│   │   ├── auth.go
│   │   └── api.go
│   └── models/
│       └── user.go
├── static/
│   ├── css/
│   ├── js/
│   └── images/
└── README.md
```

## Characteristics

- **`internal/` package**: Private application code
- **Feature-based grouping**: Components organized by type
- **Separation of layout/pages/shared**: Clear categorization
- **`cmd/` for entrypoints**: Follows Go conventions

## Complete Example: E-commerce Site

### 1. Component Organization

**internal/components/layout/base.templ:**
```templ
package layout

templ Base(title string) {
    <!DOCTYPE html>
    <html>
        <head>
            <title>{ title } - Shop</title>
            <link rel="stylesheet" href="/static/css/main.css"/>
            <script src="https://unpkg.com/htmx.org@1.9.10"></script>
        </head>
        <body>
            @Header()
            <main class="container">
                { children... }
            </main>
            @Footer()
        </body>
    </html>
}
```

**internal/components/pages/products.templ:**
```templ
package pages

import "project/internal/models"
import "project/internal/components/shared"

templ ProductList(products []models.Product) {
    <div class="products-grid">
        for _, product := range products {
            @shared.ProductCard(product)
        }
    </div>
}
```

**internal/components/shared/product-card.templ:**
```templ
package shared

import "project/internal/models"

templ ProductCard(product models.Product) {
    <div class="card">
        <img src={ product.ImageURL } alt={ product.Name }/>
        <h3>{ product.Name }</h3>
        <p class="price">${ fmt.Sprintf("%.2f", product.Price) }</p>
        @Button("Add to Cart", "/cart/add?id=" + product.ID)
    </div>
}
```

**internal/components/shared/button.templ:**
```templ
package shared

templ Button(text string, href string) {
    <a href={ templ.URL(href) } class="btn btn-primary">
        { text }
    </a>
}
```

### 2. Handler Organization

**internal/handlers/pages.go:**
```go
package handlers

import (
    "net/http"
    "project/internal/components/layout"
    "project/internal/components/pages"
)

func Home(w http.ResponseWriter, r *http.Request) {
    layout.Base("Home")(
        pages.Home(),
    ).Render(r.Context(), w)
}

func About(w http.ResponseWriter, r *http.Request) {
    layout.Base("About")(
        pages.About(),
    ).Render(r.Context(), w)
}
```

**internal/handlers/products.go:**
```go
package handlers

import (
    "net/http"
    "project/internal/components/layout"
    "project/internal/components/pages"
    "project/internal/models"
)

type ProductHandler struct {
    repo models.ProductRepository
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
    products, err := h.repo.GetAll(r.Context())
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    layout.Base("Products")(
        pages.ProductList(products),
    ).Render(r.Context(), w)
}

func (h *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    product, err := h.repo.GetByID(r.Context(), id)
    if err != nil {
        http.Error(w, "Not found", 404)
        return
    }

    layout.Base(product.Name)(
        pages.ProductDetail(product),
    ).Render(r.Context(), w)
}
```

### 3. Main Application

**cmd/server/main.go:**
```go
package main

import (
    "log"
    "net/http"
    "project/internal/handlers"
    "project/internal/models"
)

func main() {
    // Initialize dependencies
    productRepo := models.NewProductRepository()
    productHandler := &handlers.ProductHandler{repo: productRepo}

    // Static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Routes
    http.HandleFunc("/", handlers.Home)
    http.HandleFunc("/about", handlers.About)
    http.HandleFunc("/products", productHandler.List)
    http.HandleFunc("/products/detail", productHandler.Detail)

    // Start
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## Component Organization Strategies

### By Type (Recommended for Medium)

```
internal/components/
├── layout/          # Page layouts
│   ├── base.templ
│   ├── admin.templ
│   └── public.templ
├── pages/           # Full page components
│   ├── home.templ
│   ├── products.templ
│   └── dashboard.templ
└── shared/          # Reusable UI components
    ├── button.templ
    ├── card.templ
    ├── form.templ
    └── modal.templ
```

**Pros**:
- Easy to find components by type
- Clear separation of concerns
- Works well for 10-50 components

**Cons**:
- May need reorganization as project grows
- Related components may be scattered

### By Feature (Alternative)

```
internal/components/
├── auth/
│   ├── login-form.templ
│   ├── register-form.templ
│   └── password-reset.templ
├── products/
│   ├── product-list.templ
│   ├── product-card.templ
│   └── product-detail.templ
├── cart/
│   ├── cart-summary.templ
│   └── checkout-form.templ
└── shared/
    ├── layout.templ
    └── button.templ
```

**Pros**:
- Related code together
- Easier to find feature-specific components
- Scales better to large projects

**Cons**:
- Shared components need clear home
- May duplicate similar components

## Import Pattern

```go
import (
    "project/internal/components/layout"
    "project/internal/components/pages"
    "project/internal/components/shared"
)

// Usage
layout.Base("Title")(
    pages.ProductList(
        shared.ProductCard(),
        shared.Button(),
    ),
).Render(ctx, w)
```

## Handler Patterns

### Option 1: Grouped by Page

```
internal/handlers/
├── pages.go         # Home, About, Contact
├── products.go      # Product pages
├── cart.go          # Cart & Checkout
└── admin.go         # Admin pages
```

### Option 2: Handler Structs (Recommended)

```go
// internal/handlers/products.go
type ProductHandler struct {
    repo ProductRepository
    cache Cache
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) { }
func (h *ProductHandler) Detail(w http.ResponseWriter, r *http.Request) { }
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) { }
```

**Benefits**:
- Dependency injection
- Easier testing
- Clear structure

## Models & Business Logic

Keep separate from web layer:

```
internal/
├── models/
│   ├── product.go
│   ├── user.go
│   └── order.go
├── services/
│   ├── product_service.go
│   └── order_service.go
└── repositories/
    ├── product_repo.go
    └── user_repo.go
```

## Real-World Examples

### SaaS Dashboard

```
saas-app/
├── cmd/server/main.go
├── internal/
│   ├── components/
│   │   ├── layout/
│   │   │   ├── app.templ
│   │   │   └── public.templ
│   │   ├── pages/
│   │   │   ├── dashboard.templ
│   │   │   ├── settings.templ
│   │   │   └── billing.templ
│   │   └── shared/
│   │       ├── sidebar.templ
│   │       ├── stat-card.templ
│   │       └── chart.templ
│   ├── handlers/
│   │   ├── dashboard.go
│   │   ├── settings.go
│   │   └── api.go
│   └── models/
└── static/
```

### E-commerce

```
shop/
├── cmd/server/main.go
├── internal/
│   ├── components/
│   │   ├── layout/
│   │   ├── pages/
│   │   │   ├── home.templ
│   │   │   ├── products.templ
│   │   │   ├── cart.templ
│   │   │   └── checkout.templ
│   │   └── shared/
│   │       ├── product-card.templ
│   │       ├── cart-item.templ
│   │       └── payment-form.templ
│   ├── handlers/
│   ├── models/
│   └── services/
└── static/
```

### Content Management

```
cms/
├── cmd/server/main.go
├── internal/
│   ├── components/
│   │   ├── layout/
│   │   │   ├── admin.templ
│   │   │   └── public.templ
│   │   ├── pages/
│   │   │   ├── posts.templ
│   │   │   ├── editor.templ
│   │   │   └── media.templ
│   │   └── shared/
│   ├── handlers/
│   │   ├── admin/
│   │   │   ├── posts.go
│   │   │   └── media.go
│   │   └── public/
│   │       └── pages.go
│   └── models/
└── static/
```

## Migration from Small Structure

Coming from flat structure:

1. **Create `internal/` directory**
   ```bash
   mkdir -p internal/components/{layout,pages,shared}
   ```

2. **Move components by type**
   ```bash
   mv components/layout.templ internal/components/layout/base.templ
   mv components/home.templ internal/components/pages/
   mv components/button.templ internal/components/shared/
   ```

3. **Update imports**
   ```go
   // Before
   import "project/components"

   // After
   import (
       "project/internal/components/layout"
       "project/internal/components/pages"
       "project/internal/components/shared"
   )
   ```

4. **Update package declarations in .templ files**
   ```templ
   // Before
   package components

   // After
   package layout  // or pages, or shared
   ```

5. **Regenerate**
   ```bash
   templ generate
   ```

## Best Practices

1. **Use `internal/`**: Keeps implementation private
2. **Organize by type first**: layout/pages/shared for medium projects
3. **Handler structs**: Better dependency injection
4. **Separate concerns**: Web layer (components/handlers) vs business logic (models/services)
5. **Consistent naming**: Match handler files to component folders
6. **Test handlers**: Use `httptest` for handler tests

## When to Move to Large Structure

Consider [structure-large.md](./structure-large.md) when:
- **50+ components**: Need domain-driven organization
- **Multiple teams**: Different teams own different features
- **Complex business logic**: Need clear separation of layers
- **Microservices transition**: Preparing for service extraction

## Checklist

- [ ] Components organized by type (layout/pages/shared)
- [ ] Using `internal/` for private code
- [ ] Handlers use structs for dependency injection
- [ ] Models separate from web layer
- [ ] Static files organized by type
- [ ] Import paths use `internal/` prefix
- [ ] Ready to scale to large when needed
