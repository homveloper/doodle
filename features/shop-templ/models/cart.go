package models

import (
	"sync"
)

// CartItem represents a product in the shopping cart
type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Cart represents a shopping cart
type Cart struct {
	mu    sync.RWMutex
	Items []CartItem `json:"items"`
	Total float64    `json:"total"`
}

// NewCart creates a new empty cart
func NewCart() *Cart {
	return &Cart{
		Items: make([]CartItem, 0),
		Total: 0,
	}
}

// AddItem adds a product to the cart or increases quantity if it already exists
func (c *Cart) AddItem(product Product, quantity int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if product already exists in cart
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += quantity
			c.calculateTotal()
			return
		}
	}

	// Add new item
	c.Items = append(c.Items, CartItem{
		Product:  product,
		Quantity: quantity,
	})
	c.calculateTotal()
}

// UpdateQuantity updates the quantity of a product in the cart
// If quantity is 0, the item is removed
func (c *Cart) UpdateQuantity(productID int, quantity int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if quantity == 0 {
		c.removeItemUnlocked(productID)
		return
	}

	for i, item := range c.Items {
		if item.Product.ID == productID {
			c.Items[i].Quantity = quantity
			c.calculateTotal()
			return
		}
	}
}

// RemoveItem removes a product from the cart
func (c *Cart) RemoveItem(productID int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.removeItemUnlocked(productID)
}

// removeItemUnlocked removes an item without locking (internal use)
func (c *Cart) removeItemUnlocked(productID int) {
	for i, item := range c.Items {
		if item.Product.ID == productID {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			c.calculateTotal()
			return
		}
	}
}

// Clear removes all items from the cart
func (c *Cart) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Items = make([]CartItem, 0)
	c.Total = 0
}

// GetItemCount returns the total number of items in the cart
func (c *Cart) GetItemCount() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	count := 0
	for _, item := range c.Items {
		count += item.Quantity
	}
	return count
}

// calculateTotal calculates the total price of all items in the cart
// This should be called after any modification to cart items
func (c *Cart) calculateTotal() {
	total := 0.0
	for _, item := range c.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	c.Total = total
}
