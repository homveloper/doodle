package handlers

import (
	"net/http"

	"github.com/homveloper/doodle/features/shop-templ/models"
	"github.com/homveloper/doodle/features/shop-templ/templates"
)

type ProductHandler struct {
	store *models.ProductStore
	cart  *models.Cart
}

func NewProductHandler(store *models.ProductStore, cart *models.Cart) *ProductHandler {
	return &ProductHandler{
		store: store,
		cart:  cart,
	}
}

// HandleHome renders the home page with all products
func (h *ProductHandler) HandleHome(w http.ResponseWriter, r *http.Request) {
	products := h.store.GetAll()
	categories := h.store.GetCategories()

	component := templates.Layout("홈", h.cart)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render product list inside layout
	w.Header().Set("Content-Type", "text/html")
	templates.ProductList(products, categories).Render(r.Context(), w)
}

// HandleProducts returns filtered products (HTMX endpoint)
func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")

	var products []models.Product
	if category != "" {
		products = h.store.FilterByCategory(category)
	} else {
		products = h.store.GetAll()
	}

	categories := h.store.GetCategories()

	component := templates.ProductList(products, categories)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleSearch handles product search (HTMX endpoint)
func (h *ProductHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	products := h.store.Search(query)

	// Just render the product grid without categories
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<div class="product-grid">`))

	if len(products) == 0 {
		templates.EmptyState("🔍", "검색 결과가 없습니다", "다른 검색어를 시도해보세요").Render(r.Context(), w)
	} else {
		for _, product := range products {
			templates.ProductCard(product).Render(r.Context(), w)
		}
	}

	w.Write([]byte(`</div>`))
}

// HandleCategories renders the categories page
func (h *ProductHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	categories := h.store.GetCategories()

	component := templates.Layout("카테고리", h.cart)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render categories list
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<div style="padding: 20px;">`))
	w.Write([]byte(`<h2 style="margin-bottom: 16px; font-size: 24px; font-weight: 700;">카테고리</h2>`))
	w.Write([]byte(`<div style="display: flex; flex-direction: column; gap: 12px;">`))

	for _, category := range categories {
		w.Write([]byte(`<a href="/products?category=` + category + `" style="padding: 16px; background: white; border-radius: 12px; text-decoration: none; color: #333; font-weight: 600; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">`))
		w.Write([]byte(category))
		w.Write([]byte(`</a>`))
	}

	w.Write([]byte(`</div></div>`))
}
