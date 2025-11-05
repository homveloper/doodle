package handlers

import (
	"net/http"
	"strconv"

	"github.com/homveloper/doodle/features/shop-templ/models"
	"github.com/homveloper/doodle/features/shop-templ/templates"
)

type CartHandler struct {
	store *models.ProductStore
	cart  *models.Cart
}

func NewCartHandler(store *models.ProductStore, cart *models.Cart) *CartHandler {
	return &CartHandler{
		store: store,
		cart:  cart,
	}
}

// HandleCart renders the cart drawer
func (h *CartHandler) HandleCart(w http.ResponseWriter, r *http.Request) {
	component := templates.CartDrawer(h.cart)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleAddToCart adds a product to the cart
func (h *CartHandler) HandleAddToCart(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	quantityStr := r.URL.Query().Get("quantity")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	quantity := 1
	if quantityStr != "" {
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil || quantity < 1 {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}
	}

	product, exists := h.store.GetByID(productID)
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Check stock
	if product.Stock < quantity {
		http.Error(w, "Insufficient stock", http.StatusBadRequest)
		return
	}

	h.cart.AddItem(product, quantity)

	// Return updated cart badge with OOB swap
	component := templates.CartBadge(h.cart.GetItemCount())
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleUpdateCart updates the quantity of a product in the cart
func (h *CartHandler) HandleUpdateCart(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")
	quantityStr := r.URL.Query().Get("quantity")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity < 0 {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Check stock if increasing quantity
	if quantity > 0 {
		product, exists := h.store.GetByID(productID)
		if exists && product.Stock < quantity {
			http.Error(w, "Insufficient stock", http.StatusBadRequest)
			return
		}
	}

	h.cart.UpdateQuantity(productID, quantity)

	// Return updated cart drawer
	component := templates.CartDrawer(h.cart)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleRemoveFromCart removes a product from the cart
func (h *CartHandler) HandleRemoveFromCart(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("product_id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	h.cart.RemoveItem(productID)

	// Return updated cart drawer
	component := templates.CartDrawer(h.cart)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// HandleClearCart clears all items from the cart
func (h *CartHandler) HandleClearCart(w http.ResponseWriter, r *http.Request) {
	h.cart.Clear()

	// Return updated cart drawer
	component := templates.CartDrawer(h.cart)
	err := component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
