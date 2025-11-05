package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/homveloper/doodle/features/blog-templ/handlers"
	"github.com/homveloper/doodle/features/blog-templ/models"
)

func main() {
	// Create store and handler
	store := models.NewStore()
	handler := handlers.New(store)

	// Register routes
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/search", handler.Search)
	http.HandleFunc("/new", handler.NewPostForm)
	http.HandleFunc("/posts", handler.CreatePost)

	// Start server
	port := 8080
	fmt.Printf("🚀 Blog server starting on http://localhost:%d\n", port)
	fmt.Println("📝 Try searching for: templ, htmx, go, web development")
	fmt.Println("✏️  Click 'Write New Post' to create your own posts!")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
