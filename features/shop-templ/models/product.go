package models

import (
	"strings"
	"sync"
)

// Product represents an item in the e-commerce store
type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ImageURL    string   `json:"imageUrl"`
	Category    string   `json:"category"`
	Stock       int      `json:"stock"`
	Tags        []string `json:"tags"`
}

// ProductStore manages products with thread-safe operations
type ProductStore struct {
	mu       sync.RWMutex
	products map[int]Product
	nextID   int
}

// NewProductStore creates a new product store
func NewProductStore() *ProductStore {
	return &ProductStore{
		products: make(map[int]Product),
		nextID:   1,
	}
}

// Add adds a new product to the store and returns it with an assigned ID
func (s *ProductStore) Add(product Product) Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	product.ID = s.nextID
	s.nextID++
	s.products[product.ID] = product

	return product
}

// GetByID retrieves a product by its ID
func (s *ProductStore) GetByID(id int) (Product, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	product, exists := s.products[id]
	return product, exists
}

// GetAll returns all products
func (s *ProductStore) GetAll() []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	products := make([]Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}

	return products
}

// Search searches for products by name or description
func (s *ProductStore) Search(query string) []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if query == "" {
		return s.getAllUnlocked()
	}

	query = strings.ToLower(query)
	results := make([]Product, 0)

	for _, p := range s.products {
		nameMatch := strings.Contains(strings.ToLower(p.Name), query)
		descMatch := strings.Contains(strings.ToLower(p.Description), query)

		if nameMatch || descMatch {
			results = append(results, p)
		}
	}

	return results
}

// FilterByCategory returns products in a specific category
func (s *ProductStore) FilterByCategory(category string) []Product {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if category == "" {
		return s.getAllUnlocked()
	}

	results := make([]Product, 0)
	for _, p := range s.products {
		if p.Category == category {
			results = append(results, p)
		}
	}

	return results
}

// GetCategories returns a list of unique categories
func (s *ProductStore) GetCategories() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	categoryMap := make(map[string]bool)
	for _, p := range s.products {
		categoryMap[p.Category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for cat := range categoryMap {
		categories = append(categories, cat)
	}

	return categories
}

// getAllUnlocked returns all products without locking (internal use only)
func (s *ProductStore) getAllUnlocked() []Product {
	products := make([]Product, 0, len(s.products))
	for _, p := range s.products {
		products = append(products, p)
	}
	return products
}
