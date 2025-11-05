package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/homveloper/doodle/features/shop-templ/handlers"
	"github.com/homveloper/doodle/features/shop-templ/models"
	"github.com/homveloper/doodle/features/shop-templ/templates"
)

func main() {
	// Initialize store and cart
	store := models.NewProductStore()
	cart := models.NewCart()

	// Seed sample data
	seedData(store)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(store, cart)
	cartHandler := handlers.NewCartHandler(store, cart)

	// Setup routes
	mux := http.NewServeMux()

	// Product routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		products := store.GetAll()
		categories := store.GetCategories()

		component := templates.Layout("홈", cart)
		component.Render(r.Context(), w)

		// Write product list inside
		w.Write([]byte(`<div class="product-container">`))
		templates.ProductList(products, categories).Render(r.Context(), w)
		w.Write([]byte(`</div>`))
	})
	mux.HandleFunc("/products", productHandler.HandleProducts)
	mux.HandleFunc("/search", productHandler.HandleSearch)
	mux.HandleFunc("/categories", productHandler.HandleCategories)

	// Cart routes
	mux.HandleFunc("/cart", cartHandler.HandleCart)
	mux.HandleFunc("/cart/add", cartHandler.HandleAddToCart)
	mux.HandleFunc("/cart/update", cartHandler.HandleUpdateCart)
	mux.HandleFunc("/cart/remove", cartHandler.HandleRemoveFromCart)
	mux.HandleFunc("/cart/clear", cartHandler.HandleClearCart)

	// Start server
	port := ":8080"
	fmt.Printf("🛍️  Shop app running at http://localhost%s\n", port)
	fmt.Println("📱 Open in mobile viewport (430px) for best experience")
	log.Fatal(http.ListenAndServe(port, mux))
}

func seedData(store *models.ProductStore) {
	products := []models.Product{
		{
			Name:        "무선 이어폰",
			Description: "고품질 사운드와 노이즈 캔슬링 기능",
			Price:       129000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       15,
			Tags:        []string{"audio", "wireless"},
		},
		{
			Name:        "스마트워치",
			Description: "건강 추적 및 알림 기능",
			Price:       299000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       8,
			Tags:        []string{"wearable", "smart"},
		},
		{
			Name:        "백팩",
			Description: "노트북 수납 가능한 여행용 백팩",
			Price:       89000,
			ImageURL:    "",
			Category:    "패션",
			Stock:       20,
			Tags:        []string{"bag", "travel"},
		},
		{
			Name:        "텀블러",
			Description: "보온/보냉 스테인리스 텀블러",
			Price:       35000,
			ImageURL:    "",
			Category:    "생활용품",
			Stock:       50,
			Tags:        []string{"bottle", "insulated"},
		},
		{
			Name:        "USB-C 케이블",
			Description: "고속 충전 및 데이터 전송",
			Price:       19000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       100,
			Tags:        []string{"cable", "usb-c"},
		},
		{
			Name:        "무선 마우스",
			Description: "인체공학적 디자인의 무선 마우스",
			Price:       45000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       30,
			Tags:        []string{"mouse", "wireless"},
		},
		{
			Name:        "노트북 파우치",
			Description: "13인치 노트북용 보호 파우치",
			Price:       25000,
			ImageURL:    "",
			Category:    "패션",
			Stock:       25,
			Tags:        []string{"laptop", "case"},
		},
		{
			Name:        "블루투스 스피커",
			Description: "휴대용 방수 스피커",
			Price:       79000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       12,
			Tags:        []string{"speaker", "bluetooth"},
		},
		{
			Name:        "손목 보호대",
			Description: "키보드 사용 시 손목 보호",
			Price:       15000,
			ImageURL:    "",
			Category:    "생활용품",
			Stock:       40,
			Tags:        []string{"ergonomic", "wrist"},
		},
		{
			Name:        "스마트폰 거치대",
			Description: "각도 조절 가능한 거치대",
			Price:       22000,
			ImageURL:    "",
			Category:    "전자제품",
			Stock:       35,
			Tags:        []string{"phone", "stand"},
		},
		{
			Name:        "캔버스 토트백",
			Description: "친환경 에코백",
			Price:       18000,
			ImageURL:    "",
			Category:    "패션",
			Stock:       60,
			Tags:        []string{"bag", "eco"},
		},
		{
			Name:        "LED 데스크 램프",
			Description: "밝기 조절 가능 LED 램프",
			Price:       65000,
			ImageURL:    "",
			Category:    "생활용품",
			Stock:       18,
			Tags:        []string{"lamp", "led"},
		},
	}

	for _, p := range products {
		store.Add(p)
	}

	fmt.Printf("✅ Seeded %d products\n", len(products))
}
