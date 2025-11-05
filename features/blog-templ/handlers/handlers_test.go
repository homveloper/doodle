package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/homveloper/doodle/features/blog-templ/models"
)

func TestNew(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	if handler == nil {
		t.Fatal("New() returned nil")
	}

	if handler.store == nil {
		t.Error("Handler has nil store")
	}
}

func TestIndexHandler(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler.Index(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()

	// Check for essential HTML elements
	expectedElements := []string{
		"html",
		"Blog Doodle",
		"search-input",
		"hx-get",
		"post-list",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(body, elem) {
			t.Errorf("Response body missing expected element: %s", elem)
		}
	}

	// Ensure it's HTML
	if !strings.Contains(body, "<html") {
		t.Error("Response doesn't appear to be HTML")
	}
}

func TestSearchHandler(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	tests := []struct {
		name          string
		query         string
		shouldContain []string
		shouldHavePosts bool
	}{
		{
			name:          "Search for templ",
			query:         "templ",
			shouldContain: []string{"Templ", "HTMX"},
			shouldHavePosts: true,
		},
		{
			name:          "Search for htmx",
			query:         "htmx",
			shouldContain: []string{"HTMX"},
			shouldHavePosts: true,
		},
		{
			name:          "Empty search",
			query:         "",
			shouldContain: []string{"post-card"},
			shouldHavePosts: true,
		},
		{
			name:          "No results",
			query:         "xyz123nonexistent",
			shouldContain: []string{"No posts found"},
			shouldHavePosts: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/search?q="+tt.query, nil)
			w := httptest.NewRecorder()

			handler.Search(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status 200, got %d", resp.StatusCode)
			}

			body := w.Body.String()

			for _, expected := range tt.shouldContain {
				if !strings.Contains(body, expected) {
					t.Errorf("Response body missing expected content: %s", expected)
				}
			}

			if tt.shouldHavePosts {
				if !strings.Contains(body, "post-card") && !strings.Contains(body, "posts") {
					t.Error("Expected post content in response")
				}
			}
		})
	}
}

func TestSearchHandlerSpecialCharacters(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	specialQueries := []string{
		"go",
		"C++",
		"type-safe",
		"web development",
	}

	for _, query := range specialQueries {
		req := httptest.NewRequest("GET", "/search?q="+url.QueryEscape(query), nil)
		w := httptest.NewRecorder()

		handler.Search(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Query '%s' returned status %d, expected 200", query, resp.StatusCode)
		}
	}
}

func TestSearchHandlerCaseInsensitive(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	queries := []string{"templ", "TEMPL", "TeMpL"}
	var bodies []string

	for _, query := range queries {
		req := httptest.NewRequest("GET", "/search?q="+query, nil)
		w := httptest.NewRecorder()

		handler.Search(w, req)

		bodies = append(bodies, w.Body.String())
	}

	// All results should be similar (same number of posts)
	for i := 1; i < len(bodies); i++ {
		if len(bodies[i]) != len(bodies[0]) {
			t.Error("Search results differ for case variations")
		}
	}
}

func TestNewPostFormHandler(t *testing.T) {
	store := models.NewStore()
	handler := New(store)

	req := httptest.NewRequest("GET", "/new", nil)
	w := httptest.NewRecorder()

	handler.NewPostForm(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()

	// Check for form elements
	expectedElements := []string{
		"html",
		"form",
		"title",
		"content",
		"tags",
		"hx-post",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(body, elem) {
			t.Errorf("Response body missing expected element: %s", elem)
		}
	}
}

func TestCreatePostHandler(t *testing.T) {
	tests := []struct {
		name           string
		formData       url.Values
		expectedStatus int
		shouldContain  []string
	}{
		{
			name: "Valid post",
			formData: url.Values{
				"title":   {"Test Post"},
				"content": {"Test Content"},
				"tags":    {"test,go"},
			},
			expectedStatus: http.StatusOK,
			shouldContain:  []string{"Test Post", "Test Content", "post-card"},
		},
		{
			name: "Missing title",
			formData: url.Values{
				"content": {"Test Content"},
				"tags":    {"test"},
			},
			expectedStatus: http.StatusBadRequest,
			shouldContain:  []string{},
		},
		{
			name: "Missing content",
			formData: url.Values{
				"title": {"Test Post"},
				"tags":  {"test"},
			},
			expectedStatus: http.StatusBadRequest,
			shouldContain:  []string{},
		},
		{
			name: "No tags",
			formData: url.Values{
				"title":   {"Test Post"},
				"content": {"Test Content"},
			},
			expectedStatus: http.StatusOK,
			shouldContain:  []string{"Test Post", "Test Content"},
		},
		{
			name: "Multiple tags",
			formData: url.Values{
				"title":   {"Test Post"},
				"content": {"Test Content"},
				"tags":    {"go,htmx,tutorial"},
			},
			expectedStatus: http.StatusOK,
			shouldContain:  []string{"Test Post", "go", "htmx", "tutorial"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := models.NewStore()
			handler := New(store)

			req := httptest.NewRequest("POST", "/posts", strings.NewReader(tt.formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()

			handler.CreatePost(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				body := w.Body.String()
				for _, expected := range tt.shouldContain {
					if !strings.Contains(body, expected) {
						t.Errorf("Response body missing expected content: %s", expected)
					}
				}

				// Verify post was added to store
				posts := store.GetAll()
				if posts[0].Title != tt.formData.Get("title") {
					t.Errorf("Expected first post title '%s', got '%s'", tt.formData.Get("title"), posts[0].Title)
				}
			}
		})
	}
}
