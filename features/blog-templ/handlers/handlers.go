package handlers

import (
	"net/http"
	"strings"

	"github.com/homveloper/doodle/features/blog-templ/models"
	"github.com/homveloper/doodle/features/blog-templ/templates"
)

// Handler manages HTTP requests for the blog
type Handler struct {
	store *models.Store
}

// New creates a new handler with a post store
func New(store *models.Store) *Handler {
	return &Handler{
		store: store,
	}
}

// Index handles the home page
func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	posts := h.store.GetAll()
	templates.Index(posts).Render(r.Context(), w)
}

// Search handles the search endpoint
func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	posts := h.store.Search(query)
	templates.PostList(posts).Render(r.Context(), w)
}

// NewPostForm handles the new post form page
func (h *Handler) NewPostForm(w http.ResponseWriter, r *http.Request) {
	templates.NewPostForm().Render(r.Context(), w)
}

// CreatePost handles post creation
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	title := strings.TrimSpace(r.FormValue("title"))
	content := strings.TrimSpace(r.FormValue("content"))
	tagsStr := strings.TrimSpace(r.FormValue("tags"))

	// Parse tags (comma-separated)
	var tags []string
	if tagsStr != "" {
		for _, tag := range strings.Split(tagsStr, ",") {
			trimmed := strings.TrimSpace(tag)
			if trimmed != "" {
				tags = append(tags, trimmed)
			}
		}
	}

	// Create post with default author
	post := models.Post{
		Title:   title,
		Content: content,
		Author:  "Blog Author", // Default author as per user preference
		Tags:    tags,
	}

	// Add post to store
	if err := h.store.Add(post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the newly added post (it's at the beginning)
	posts := h.store.GetAll()
	newPost := posts[0]

	// Return the new post card for HTMX to insert
	templates.PostCard(newPost).Render(r.Context(), w)
}
