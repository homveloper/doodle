# Large Project Structure (50+ components)

For enterprise applications, multi-tenant SaaS, complex platforms, or microservice-ready architectures.

## Recommended Structure

```
project/
├── go.mod
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── app/                    # Business logic layer
│   │   ├── models/
│   │   ├── services/
│   │   └── repositories/
│   ├── web/                    # Web/presentation layer
│   │   ├── components/
│   │   │   ├── auth/           # By domain/feature
│   │   │   │   ├── login.templ
│   │   │   │   ├── register.templ
│   │   │   │   └── forgot-password.templ
│   │   │   ├── dashboard/
│   │   │   │   ├── overview.templ
│   │   │   │   ├── stats.templ
│   │   │   │   └── charts.templ
│   │   │   ├── products/
│   │   │   │   ├── list.templ
│   │   │   │   ├── detail.templ
│   │   │   │   ├── edit.templ
│   │   │   │   └── create.templ
│   │   │   └── shared/
│   │   │       ├── layouts/
│   │   │       │   ├── app.templ
│   │   │       │   ├── public.templ
│   │   │       │   └── admin.templ
│   │   │       └── ui/
│   │   │           ├── button.templ
│   │   │           ├── input.templ
│   │   │           ├── table.templ
│   │   │           └── modal.templ
│   │   ├── handlers/
│   │   │   ├── auth/
│   │   │   │   └── handlers.go
│   │   │   ├── dashboard/
│   │   │   │   └── handlers.go
│   │   │   └── products/
│   │   │       └── handlers.go
│   │   └── middleware/
│   │       ├── auth.go
│   │       ├── logging.go
│   │       └── cors.go
│   └── config/
│       └── config.go
├── pkg/                        # Public packages (reusable)
│   ├── validator/
│   └── pagination/
├── web/                        # Frontend assets
│   └── static/
│       ├── css/
│       ├── js/
│       └── images/
├── migrations/                 # Database migrations
├── scripts/                    # Build/deploy scripts
└── deployments/                # Deployment configs
    ├── docker/
    └── kubernetes/
```

## Characteristics

- **Domain-driven design**: Components organized by business domain
- **Clear layering**: app (business) vs web (presentation)
- **Team scalability**: Different teams can own different domains
- **Microservice-ready**: Easy to extract domains into services
- **Public packages**: Reusable code in `pkg/`

## Domain-Based Component Organization

### Feature Domains

```
internal/web/components/
├── auth/                       # Authentication & authorization
│   ├── login-form.templ
│   ├── register-form.templ
│   ├── forgot-password.templ
│   ├── reset-password.templ
│   └── oauth-buttons.templ
├── users/                      # User management
│   ├── user-list.templ
│   ├── user-profile.templ
│   ├── user-settings.templ
│   └── user-avatar.templ
├── products/                   # Product catalog
│   ├── product-list.templ
│   ├── product-grid.templ
│   ├── product-card.templ
│   ├── product-detail.templ
│   ├── product-reviews.templ
│   └── product-editor.templ
├── orders/                     # Order management
│   ├── order-list.templ
│   ├── order-detail.templ
│   ├── order-status.templ
│   └── order-tracking.templ
├── billing/                    # Billing & payments
│   ├── invoices.templ
│   ├── payment-methods.templ
│   └── subscription.templ
└── shared/                     # Shared components
    ├── layouts/
    └── ui/
```

### Shared Components Structure

```
internal/web/components/shared/
├── layouts/
│   ├── app.templ              # Authenticated app layout
│   ├── public.templ           # Public pages layout
│   ├── admin.templ            # Admin panel layout
│   └── email.templ            # Email template layout
└── ui/                        # Design system components
    ├── forms/
    │   ├── input.templ
    │   ├── textarea.templ
    │   ├── select.templ
    │   └── checkbox.templ
    ├── buttons/
    │   ├── button.templ
    │   ├── icon-button.templ
    │   └── link-button.templ
    ├── feedback/
    │   ├── alert.templ
    │   ├── toast.templ
    │   └── loading.templ
    └── data/
        ├── table.templ
        ├── pagination.templ
        └── empty-state.templ
```

## Complete Example: Multi-tenant SaaS

### 1. Authentication Domain

**internal/web/components/auth/login-form.templ:**
```templ
package auth

import "project/internal/web/components/shared/ui/forms"
import "project/internal/web/components/shared/ui/buttons"

templ LoginForm(csrf string) {
    <form class="auth-form" hx-post="/auth/login" hx-target="#auth-container">
        <input type="hidden" name="csrf" value={ csrf }/>

        @forms.Input(forms.InputProps{
            Name: "email",
            Type: "email",
            Label: "Email",
            Required: true,
        })

        @forms.Input(forms.InputProps{
            Name: "password",
            Type: "password",
            Label: "Password",
            Required: true,
        })

        <div class="form-actions">
            @buttons.Button(buttons.ButtonProps{
                Type: "submit",
                Text: "Sign In",
                Variant: "primary",
            })
            <a href="/auth/forgot-password">Forgot password?</a>
        </div>
    </form>
}
```

### 2. Dashboard Domain

**internal/web/components/dashboard/overview.templ:**
```templ
package dashboard

import "project/internal/app/models"
import "project/internal/web/components/shared/ui/data"

templ Overview(stats models.DashboardStats) {
    <div class="dashboard-overview">
        <div class="stats-grid">
            @StatCard("Revenue", fmt.Sprintf("$%.2f", stats.Revenue), "+12%")
            @StatCard("Users", fmt.Sprint(stats.Users), "+5%")
            @StatCard("Orders", fmt.Sprint(stats.Orders), "+8%")
        </div>

        <div class="recent-orders">
            <h2>Recent Orders</h2>
            @data.Table(data.TableProps{
                Headers: []string{"ID", "Customer", "Amount", "Status"},
                Rows: formatOrders(stats.RecentOrders),
            })
        </div>
    </div>
}
```

### 3. Handler Organization

**internal/web/handlers/auth/handlers.go:**
```go
package auth

import (
    "net/http"
    "project/internal/app/services"
    "project/internal/web/components/auth"
    "project/internal/web/components/shared/layouts"
)

type Handler struct {
    authService *services.AuthService
    csrf        CSRFTokenGenerator
}

func NewHandler(authService *services.AuthService, csrf CSRFTokenGenerator) *Handler {
    return &Handler{
        authService: authService,
        csrf:        csrf,
    }
}

func (h *Handler) ShowLogin(w http.ResponseWriter, r *http.Request) {
    csrf := h.csrf.Generate(r)

    layouts.Public("Login")(
        auth.LoginForm(csrf),
    ).Render(r.Context(), w)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    user, err := h.authService.Authenticate(
        r.Context(),
        r.FormValue("email"),
        r.FormValue("password"),
    )

    if err != nil {
        // Return error component
        auth.LoginError(err.Error()).Render(r.Context(), w)
        return
    }

    // Set session, redirect
    h.setSession(w, r, user)
    w.Header().Set("HX-Redirect", "/dashboard")
    w.WriteHeader(http.StatusOK)
}
```

**internal/web/handlers/products/handlers.go:**
```go
package products

import (
    "net/http"
    "project/internal/app/services"
    "project/internal/web/components/products"
    "project/internal/web/components/shared/layouts"
)

type Handler struct {
    productService *services.ProductService
}

func NewHandler(productService *services.ProductService) *Handler {
    return &Handler{productService: productService}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    products, err := h.productService.GetAll(ctx)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    layouts.App("Products")(
        products.List(products),
    ).Render(ctx, w)
}

func (h *Handler) Detail(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    id := r.URL.Query().Get("id")

    product, err := h.productService.GetByID(ctx, id)
    if err != nil {
        http.Error(w, "Not found", 404)
        return
    }

    layouts.App(product.Name)(
        products.Detail(product),
    ).Render(ctx, w)
}
```

### 4. Business Logic Layer

**internal/app/services/product_service.go:**
```go
package services

import (
    "context"
    "project/internal/app/models"
    "project/internal/app/repositories"
)

type ProductService struct {
    repo repositories.ProductRepository
    cache CacheService
}

func NewProductService(repo repositories.ProductRepository, cache CacheService) *ProductService {
    return &ProductService{repo: repo, cache: cache}
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
    // Check cache
    if cached, ok := s.cache.Get("products:all"); ok {
        return cached.([]models.Product), nil
    }

    // Query database
    products, err := s.repo.FindAll(ctx)
    if err != nil {
        return nil, err
    }

    // Cache results
    s.cache.Set("products:all", products, 5*time.Minute)

    return products, nil
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
    // Business logic, validation, caching...
    return s.repo.FindByID(ctx, id)
}
```

### 5. Application Initialization

**cmd/server/main.go:**
```go
package main

import (
    "log"
    "net/http"

    "project/internal/app/repositories"
    "project/internal/app/services"
    "project/internal/web/handlers/auth"
    "project/internal/web/handlers/dashboard"
    "project/internal/web/handlers/products"
    "project/internal/web/middleware"
    "project/internal/config"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Initialize repositories
    db := initDatabase(cfg)
    productRepo := repositories.NewProductRepository(db)
    userRepo := repositories.NewUserRepository(db)

    // Initialize services
    cache := initCache(cfg)
    productService := services.NewProductService(productRepo, cache)
    authService := services.NewAuthService(userRepo)

    // Initialize handlers
    authHandler := auth.NewHandler(authService, csrf)
    dashboardHandler := dashboard.NewHandler(productService)
    productHandler := products.NewHandler(productService)

    // Setup router
    mux := http.NewServeMux()

    // Static files
    mux.Handle("/static/", http.StripPrefix("/static/",
        http.FileServer(http.Dir("web/static"))))

    // Public routes
    mux.HandleFunc("/auth/login", authHandler.ShowLogin)
    mux.HandleFunc("/auth/register", authHandler.ShowRegister)

    // Protected routes (with auth middleware)
    protected := middleware.Chain(
        middleware.Auth(authService),
        middleware.Logging(),
    )

    mux.Handle("/dashboard", protected(http.HandlerFunc(dashboardHandler.Show)))
    mux.Handle("/products", protected(http.HandlerFunc(productHandler.List)))
    mux.Handle("/products/detail", protected(http.HandlerFunc(productHandler.Detail)))

    // Start server
    log.Printf("Server starting on %s", cfg.Address)
    log.Fatal(http.ListenAndServe(cfg.Address, mux))
}
```

## Middleware Pattern

**internal/web/middleware/auth.go:**
```go
package middleware

import (
    "context"
    "net/http"
    "project/internal/app/services"
)

type contextKey string

const userKey contextKey = "user"

func Auth(authService *services.AuthService) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Get session/token
            session := getSession(r)
            if session == "" {
                http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
                return
            }

            // Validate session
            user, err := authService.ValidateSession(r.Context(), session)
            if err != nil {
                http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
                return
            }

            // Add user to context
            ctx := context.WithValue(r.Context(), userKey, user)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

func Chain(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
    return func(final http.Handler) http.Handler {
        for i := len(middlewares) - 1; i >= 0; i-- {
            final = middlewares[i](final)
        }
        return final
    }
}
```

## Public Packages (`pkg/`)

Code that could be extracted to separate libraries:

```
pkg/
├── validator/
│   ├── validator.go
│   └── rules.go
├── pagination/
│   ├── paginator.go
│   └── cursor.go
├── slug/
│   └── slug.go
└── errors/
    └── errors.go
```

## Import Pattern

```go
import (
    // App layer
    "project/internal/app/models"
    "project/internal/app/services"

    // Web layer
    "project/internal/web/components/products"
    "project/internal/web/components/shared/layouts"
    "project/internal/web/components/shared/ui/buttons"

    // Public packages
    "project/pkg/validator"
    "project/pkg/pagination"
)
```

## Real-World Examples

### Multi-tenant SaaS

```
saas-platform/
├── cmd/server/
├── internal/
│   ├── app/
│   │   ├── models/
│   │   ├── services/
│   │   └── repositories/
│   └── web/
│       ├── components/
│       │   ├── tenants/
│       │   ├── billing/
│       │   ├── analytics/
│       │   └── shared/
│       └── handlers/
└── pkg/
    ├── multitenancy/
    └── billing/
```

### Enterprise CRM

```
crm/
├── cmd/
│   ├── server/
│   ├── worker/
│   └── cli/
├── internal/
│   ├── app/
│   │   ├── contacts/
│   │   ├── deals/
│   │   ├── tasks/
│   │   └── shared/
│   └── web/
│       ├── components/
│       │   ├── contacts/
│       │   ├── deals/
│       │   ├── pipeline/
│       │   └── shared/
│       └── handlers/
└── pkg/
```

## Migration from Medium Structure

1. **Create domain directories**
   ```bash
   mkdir -p internal/web/components/{auth,dashboard,products}
   ```

2. **Move related components**
   ```bash
   # Group by domain instead of type
   mv internal/components/pages/login.templ internal/web/components/auth/
   mv internal/components/pages/dashboard.templ internal/web/components/dashboard/
   ```

3. **Extract business logic**
   ```bash
   mkdir -p internal/app/{models,services,repositories}
   # Move business logic out of handlers
   ```

4. **Update imports**
   ```go
   // From type-based
   import "project/internal/components/pages"

   // To domain-based
   import "project/internal/web/components/dashboard"
   ```

## Team Organization

Different teams can own different domains:

- **Auth Team**: `internal/web/components/auth/`, `internal/app/services/auth_service.go`
- **Products Team**: `internal/web/components/products/`, `internal/app/services/product_service.go`
- **Billing Team**: `internal/web/components/billing/`, `internal/app/services/billing_service.go`

Each team can work independently with minimal conflicts.

## Best Practices

1. **Domain-driven**: Organize by business domain, not technical type
2. **Clear layers**: Separate web (presentation) from app (business logic)
3. **Handler structs with DI**: Easy testing, clear dependencies
4. **Middleware chain**: Composable request processing
5. **Public packages**: Extract reusable code to `pkg/`
6. **Team ownership**: Each domain has a clear owner team

## Microservice Extraction

When ready to extract a domain into a microservice:

1. Domain is already isolated (`internal/web/components/billing/`, `internal/app/services/billing_service.go`)
2. Extract to separate repo
3. Expose as API
4. Update main app to call API instead of internal service

## Checklist

- [ ] Components organized by business domain
- [ ] Clear app (business) vs web (presentation) layers
- [ ] Handler structs with dependency injection
- [ ] Middleware for cross-cutting concerns
- [ ] Public packages in `pkg/` for reusable code
- [ ] Team ownership of domains
- [ ] Ready for microservice extraction if needed
