package models

import (
	"testing"
)

func TestNewCart(t *testing.T) {
	cart := NewCart()
	if cart == nil {
		t.Fatal("NewCart should return a non-nil cart")
	}
	if len(cart.Items) != 0 {
		t.Errorf("New cart should be empty, got %d items", len(cart.Items))
	}
	if cart.Total != 0 {
		t.Errorf("New cart total should be 0, got %.2f", cart.Total)
	}
}

func TestAddToCart(t *testing.T) {
	cart := NewCart()
	product := Product{ID: 1, Name: "Test Product", Price: 19.99, Stock: 10}

	cart.AddItem(product, 2)

	if len(cart.Items) != 1 {
		t.Errorf("Expected 1 item in cart, got %d", len(cart.Items))
	}

	item := cart.Items[0]
	if item.Product.ID != 1 {
		t.Errorf("Expected product ID 1, got %d", item.Product.ID)
	}
	if item.Quantity != 2 {
		t.Errorf("Expected quantity 2, got %d", item.Quantity)
	}
	if cart.Total != 39.98 {
		t.Errorf("Expected total 39.98, got %.2f", cart.Total)
	}
}

func TestAddExistingProductToCart(t *testing.T) {
	cart := NewCart()
	product := Product{ID: 1, Name: "Test Product", Price: 10.00, Stock: 10}

	cart.AddItem(product, 2)
	cart.AddItem(product, 3)

	if len(cart.Items) != 1 {
		t.Errorf("Expected 1 item in cart (same product), got %d", len(cart.Items))
	}

	if cart.Items[0].Quantity != 5 {
		t.Errorf("Expected quantity 5, got %d", cart.Items[0].Quantity)
	}

	if cart.Total != 50.00 {
		t.Errorf("Expected total 50.00, got %.2f", cart.Total)
	}
}

func TestUpdateQuantity(t *testing.T) {
	cart := NewCart()
	product := Product{ID: 1, Name: "Test Product", Price: 15.00, Stock: 10}

	cart.AddItem(product, 2)
	cart.UpdateQuantity(1, 5)

	if cart.Items[0].Quantity != 5 {
		t.Errorf("Expected quantity 5, got %d", cart.Items[0].Quantity)
	}
	if cart.Total != 75.00 {
		t.Errorf("Expected total 75.00, got %.2f", cart.Total)
	}
}

func TestUpdateQuantityToZero(t *testing.T) {
	cart := NewCart()
	product := Product{ID: 1, Name: "Test Product", Price: 10.00, Stock: 10}

	cart.AddItem(product, 2)
	cart.UpdateQuantity(1, 0)

	if len(cart.Items) != 0 {
		t.Errorf("Expected 0 items after setting quantity to 0, got %d", len(cart.Items))
	}
	if cart.Total != 0 {
		t.Errorf("Expected total 0, got %.2f", cart.Total)
	}
}

func TestRemoveFromCart(t *testing.T) {
	cart := NewCart()
	product1 := Product{ID: 1, Name: "Product 1", Price: 10.00, Stock: 10}
	product2 := Product{ID: 2, Name: "Product 2", Price: 20.00, Stock: 5}

	cart.AddItem(product1, 1)
	cart.AddItem(product2, 2)

	if len(cart.Items) != 2 {
		t.Fatalf("Expected 2 items, got %d", len(cart.Items))
	}

	cart.RemoveItem(1)

	if len(cart.Items) != 1 {
		t.Errorf("Expected 1 item after removal, got %d", len(cart.Items))
	}
	if cart.Items[0].Product.ID != 2 {
		t.Errorf("Expected remaining product ID 2, got %d", cart.Items[0].Product.ID)
	}
	if cart.Total != 40.00 {
		t.Errorf("Expected total 40.00, got %.2f", cart.Total)
	}
}

func TestRemoveNonExistentProduct(t *testing.T) {
	cart := NewCart()
	product := Product{ID: 1, Name: "Test Product", Price: 10.00, Stock: 10}

	cart.AddItem(product, 1)
	cart.RemoveItem(999) // Non-existent ID

	if len(cart.Items) != 1 {
		t.Errorf("Removing non-existent product should not affect cart, got %d items", len(cart.Items))
	}
}

func TestClearCart(t *testing.T) {
	cart := NewCart()
	product1 := Product{ID: 1, Name: "Product 1", Price: 10.00, Stock: 10}
	product2 := Product{ID: 2, Name: "Product 2", Price: 20.00, Stock: 5}

	cart.AddItem(product1, 2)
	cart.AddItem(product2, 3)

	cart.Clear()

	if len(cart.Items) != 0 {
		t.Errorf("Expected 0 items after clear, got %d", len(cart.Items))
	}
	if cart.Total != 0 {
		t.Errorf("Expected total 0 after clear, got %.2f", cart.Total)
	}
}

func TestGetItemCount(t *testing.T) {
	cart := NewCart()
	product1 := Product{ID: 1, Name: "Product 1", Price: 10.00, Stock: 10}
	product2 := Product{ID: 2, Name: "Product 2", Price: 20.00, Stock: 5}

	cart.AddItem(product1, 3)
	cart.AddItem(product2, 2)

	count := cart.GetItemCount()
	if count != 5 {
		t.Errorf("Expected item count 5, got %d", count)
	}
}

func TestCalculateTotal(t *testing.T) {
	cart := NewCart()

	// Test with multiple products
	cart.AddItem(Product{ID: 1, Price: 10.50, Stock: 10}, 2) // 21.00
	cart.AddItem(Product{ID: 2, Price: 15.99, Stock: 5}, 1)  // 15.99
	cart.AddItem(Product{ID: 3, Price: 7.25, Stock: 20}, 3)  // 21.75

	expected := 58.74
	if cart.Total != expected {
		t.Errorf("Expected total %.2f, got %.2f", expected, cart.Total)
	}
}
