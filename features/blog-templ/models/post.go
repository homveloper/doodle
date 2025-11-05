package models

import (
	"errors"
	"strings"
	"sync"
	"time"
)

// Post represents a blog post
type Post struct {
	ID        int
	Title     string
	Content   string
	Author    string
	CreatedAt time.Time
	Tags      []string
}

// Store manages blog posts
type Store struct {
	posts  []Post
	mu     sync.Mutex
	nextID int
}

// NewStore creates a new post store with sample data
func NewStore() *Store {
	return &Store{
		posts: []Post{
			{
				ID:        1,
				Title:     "Getting Started with Templ and HTMX",
				Content:   "Templ is a templating language for Go that generates type-safe HTML. Combined with HTMX, you can build dynamic web applications without writing JavaScript.",
				Author:    "Jane Doe",
				CreatedAt: time.Now().AddDate(0, 0, -7),
				Tags:      []string{"templ", "htmx", "go", "tutorial"},
			},
			{
				ID:        2,
				Title:     "Building Real-time Search with HTMX",
				Content:   "HTMX makes it easy to add AJAX requests directly in HTML. With hx-get and hx-trigger, you can create real-time search without complex JavaScript.",
				Author:    "John Smith",
				CreatedAt: time.Now().AddDate(0, 0, -5),
				Tags:      []string{"htmx", "search", "web development"},
			},
			{
				ID:        3,
				Title:     "Why Go is Great for Web Development",
				Content:   "Go's simplicity, performance, and built-in concurrency make it an excellent choice for web applications. The standard library is powerful and well-designed.",
				Author:    "Jane Doe",
				CreatedAt: time.Now().AddDate(0, 0, -3),
				Tags:      []string{"go", "web development", "backend"},
			},
			{
				ID:        4,
				Title:     "Type-Safe HTML Templates",
				Content:   "Templ generates Go code from templates, giving you compile-time safety and IDE support. No more runtime template errors!",
				Author:    "John Smith",
				CreatedAt: time.Now().AddDate(0, 0, -1),
				Tags:      []string{"templ", "go", "type safety"},
			},
		},
		nextID: 5, // Start from 5 since we have 4 sample posts
	}
}

// GetAll returns all posts
func (s *Store) GetAll() []Post {
	return s.posts
}

// Search returns posts matching the query
func (s *Store) Search(query string) []Post {
	if query == "" {
		return s.posts
	}

	query = strings.ToLower(query)
	var results []Post

	for _, post := range s.posts {
		if s.matches(post, query) {
			results = append(results, post)
		}
	}

	return results
}

// matches checks if a post matches the search query
func (s *Store) matches(post Post, query string) bool {
	// Search in title
	if strings.Contains(strings.ToLower(post.Title), query) {
		return true
	}

	// Search in content
	if strings.Contains(strings.ToLower(post.Content), query) {
		return true
	}

	// Search in author
	if strings.Contains(strings.ToLower(post.Author), query) {
		return true
	}

	// Search in tags
	for _, tag := range post.Tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}

	return false
}

// Add adds a new post to the store
func (s *Store) Add(post Post) error {
	// Validate input
	if strings.TrimSpace(post.Title) == "" {
		return errors.New("title is required")
	}
	if strings.TrimSpace(post.Content) == "" {
		return errors.New("content is required")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Set auto-generated fields
	post.ID = s.nextID
	s.nextID++
	post.CreatedAt = time.Now()

	// Add to the beginning (most recent first)
	s.posts = append([]Post{post}, s.posts...)

	return nil
}
