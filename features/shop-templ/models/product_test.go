package models

import (
	"testing"
)

func TestNewProductStore(t *testing.T) {
	store := NewProductStore()
	if store == nil {
		t.Fatal("NewProductStore should return a non-nil store")
	}
	if len(store.GetAll()) != 0 {
		t.Errorf("New store should be empty, got %d products", len(store.GetAll()))
	}
}

func TestAddProduct(t *testing.T) {
	store := NewProductStore()

	product := Product{
		Name:        "Test Product",
		Description: "A test product",
		Price:       19.99,
		ImageURL:    "/images/test.jpg",
		Category:    "Electronics",
		Stock:       10,
		Tags:        []string{"test", "sample"},
	}

	added := store.Add(product)
	if added.ID == 0 {
		t.Error("Added product should have a non-zero ID")
	}
	if added.Name != product.Name {
		t.Errorf("Expected name %s, got %s", product.Name, added.Name)
	}
}

func TestGetByID(t *testing.T) {
	store := NewProductStore()
	product := Product{Name: "Test", Price: 9.99, Category: "Test", Stock: 5}
	added := store.Add(product)

	found, exists := store.GetByID(added.ID)
	if !exists {
		t.Fatal("Product should exist")
	}
	if found.ID != added.ID {
		t.Errorf("Expected ID %d, got %d", added.ID, found.ID)
	}

	_, exists = store.GetByID(999)
	if exists {
		t.Error("Non-existent product should not be found")
	}
}

func TestGetAll(t *testing.T) {
	store := NewProductStore()

	store.Add(Product{Name: "Product 1", Price: 10.00, Category: "Cat1", Stock: 5})
	store.Add(Product{Name: "Product 2", Price: 20.00, Category: "Cat2", Stock: 3})
	store.Add(Product{Name: "Product 3", Price: 30.00, Category: "Cat1", Stock: 8})

	all := store.GetAll()
	if len(all) != 3 {
		t.Errorf("Expected 3 products, got %d", len(all))
	}
}

func TestSearchProducts(t *testing.T) {
	store := NewProductStore()

	store.Add(Product{Name: "Laptop Computer", Description: "High performance laptop", Price: 999.00, Category: "Electronics", Stock: 5})
	store.Add(Product{Name: "Wireless Mouse", Description: "Ergonomic mouse", Price: 29.99, Category: "Electronics", Stock: 20})
	store.Add(Product{Name: "Coffee Mug", Description: "Ceramic mug", Price: 12.99, Category: "Home", Stock: 50})

	tests := []struct {
		query    string
		expected int
	}{
		{"laptop", 1},
		{"mouse", 1},
		{"wireless", 1},
		{"electronics", 0}, // Search doesn't include category
		{"mug", 1},
		{"", 3}, // Empty query returns all
		{"xyz", 0},
	}

	for _, tt := range tests {
		results := store.Search(tt.query)
		if len(results) != tt.expected {
			t.Errorf("Search(%q): expected %d results, got %d", tt.query, tt.expected, len(results))
		}
	}
}

func TestFilterByCategory(t *testing.T) {
	store := NewProductStore()

	store.Add(Product{Name: "Laptop", Price: 999.00, Category: "Electronics", Stock: 5})
	store.Add(Product{Name: "Mouse", Price: 29.99, Category: "Electronics", Stock: 20})
	store.Add(Product{Name: "Mug", Price: 12.99, Category: "Home", Stock: 50})
	store.Add(Product{Name: "Shirt", Price: 24.99, Category: "Clothing", Stock: 30})

	tests := []struct {
		category string
		expected int
	}{
		{"Electronics", 2},
		{"Home", 1},
		{"Clothing", 1},
		{"Books", 0},
		{"", 4}, // Empty category returns all
	}

	for _, tt := range tests {
		results := store.FilterByCategory(tt.category)
		if len(results) != tt.expected {
			t.Errorf("FilterByCategory(%q): expected %d results, got %d", tt.category, tt.expected, len(results))
		}
	}
}

func TestGetCategories(t *testing.T) {
	store := NewProductStore()

	store.Add(Product{Name: "P1", Price: 10.00, Category: "Electronics", Stock: 5})
	store.Add(Product{Name: "P2", Price: 20.00, Category: "Electronics", Stock: 3})
	store.Add(Product{Name: "P3", Price: 30.00, Category: "Home", Stock: 8})
	store.Add(Product{Name: "P4", Price: 40.00, Category: "Clothing", Stock: 2})

	categories := store.GetCategories()
	if len(categories) != 3 {
		t.Errorf("Expected 3 unique categories, got %d", len(categories))
	}

	// Check that categories are unique
	categoryMap := make(map[string]bool)
	for _, cat := range categories {
		if categoryMap[cat] {
			t.Errorf("Duplicate category found: %s", cat)
		}
		categoryMap[cat] = true
	}
}
