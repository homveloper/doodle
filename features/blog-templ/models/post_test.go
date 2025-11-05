package models

import (
	"testing"
)

func TestNewStore(t *testing.T) {
	store := NewStore()
	if store == nil {
		t.Fatal("NewStore() returned nil")
	}

	posts := store.GetAll()
	if len(posts) == 0 {
		t.Error("Expected sample posts, got empty store")
	}
}

func TestGetAll(t *testing.T) {
	store := NewStore()
	posts := store.GetAll()

	if len(posts) != 4 {
		t.Errorf("Expected 4 posts, got %d", len(posts))
	}

	// Verify post structure
	for i, post := range posts {
		if post.ID == 0 {
			t.Errorf("Post %d has invalid ID", i)
		}
		if post.Title == "" {
			t.Errorf("Post %d has empty title", i)
		}
		if post.Author == "" {
			t.Errorf("Post %d has empty author", i)
		}
	}
}

func TestSearchEmptyQuery(t *testing.T) {
	store := NewStore()
	results := store.Search("")

	allPosts := store.GetAll()
	if len(results) != len(allPosts) {
		t.Errorf("Empty query should return all posts. Expected %d, got %d", len(allPosts), len(results))
	}
}

func TestSearchByTitle(t *testing.T) {
	store := NewStore()
	results := store.Search("templ")

	if len(results) == 0 {
		t.Error("Expected results for 'templ' query")
	}

	// Verify all results contain "templ" in title, content, author, or tags
	for _, post := range results {
		found := false
		searchQuery := "templ"

		if contains(post.Title, searchQuery) ||
			contains(post.Content, searchQuery) ||
			contains(post.Author, searchQuery) {
			found = true
		}

		for _, tag := range post.Tags {
			if contains(tag, searchQuery) {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Post '%s' doesn't match search query 'templ'", post.Title)
		}
	}
}

func TestSearchByContent(t *testing.T) {
	store := NewStore()
	results := store.Search("javascript")

	if len(results) == 0 {
		t.Error("Expected results for 'javascript' query")
	}
}

func TestSearchByAuthor(t *testing.T) {
	store := NewStore()
	results := store.Search("jane")

	if len(results) == 0 {
		t.Error("Expected results for 'jane' query")
	}

	for _, post := range results {
		if !contains(post.Author, "jane") {
			t.Errorf("Post by '%s' shouldn't be in results for 'jane' query", post.Author)
		}
	}
}

func TestSearchByTag(t *testing.T) {
	store := NewStore()
	results := store.Search("htmx")

	if len(results) == 0 {
		t.Error("Expected results for 'htmx' query")
	}

	for _, post := range results {
		hasTag := false
		for _, tag := range post.Tags {
			if contains(tag, "htmx") {
				hasTag = true
				break
			}
		}

		if !hasTag && !contains(post.Title, "htmx") && !contains(post.Content, "htmx") {
			t.Errorf("Post '%s' doesn't have 'htmx' tag or mention", post.Title)
		}
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	store := NewStore()

	lowerResults := store.Search("templ")
	upperResults := store.Search("TEMPL")
	mixedResults := store.Search("TeMpL")

	if len(lowerResults) != len(upperResults) || len(lowerResults) != len(mixedResults) {
		t.Error("Search should be case insensitive")
	}
}

func TestSearchNoResults(t *testing.T) {
	store := NewStore()
	results := store.Search("xyz123nonexistent")

	if len(results) != 0 {
		t.Errorf("Expected no results for non-existent query, got %d", len(results))
	}
}

func TestSearchMultipleMatches(t *testing.T) {
	store := NewStore()
	results := store.Search("go")

	// "go" should match multiple posts (in tags, content, etc.)
	if len(results) < 2 {
		t.Errorf("Expected multiple results for 'go' query, got %d", len(results))
	}
}

// Helper function for case-insensitive string comparison
func contains(s, substr string) bool {
	s = toLower(s)
	substr = toLower(substr)

	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func toLower(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			result[i] = r + 32
		} else {
			result[i] = r
		}
	}
	return string(result)
}

func TestAddPost(t *testing.T) {
	store := NewStore()
	initialCount := len(store.GetAll())

	newPost := Post{
		Title:   "Test Post",
		Content: "Test Content",
		Tags:    []string{"test", "go"},
	}

	err := store.Add(newPost)
	if err != nil {
		t.Fatalf("Add() failed: %v", err)
	}

	posts := store.GetAll()
	if len(posts) != initialCount+1 {
		t.Errorf("Expected %d posts, got %d", initialCount+1, len(posts))
	}

	// Verify the new post is at the beginning (most recent first)
	addedPost := posts[0]
	if addedPost.Title != newPost.Title {
		t.Errorf("Expected title '%s', got '%s'", newPost.Title, addedPost.Title)
	}
	if addedPost.Content != newPost.Content {
		t.Errorf("Expected content '%s', got '%s'", newPost.Content, addedPost.Content)
	}
	if addedPost.ID == 0 {
		t.Error("Expected auto-generated ID, got 0")
	}
	if addedPost.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}
}

func TestAddPostValidation(t *testing.T) {
	store := NewStore()

	tests := []struct {
		name    string
		post    Post
		wantErr bool
	}{
		{
			name:    "empty title",
			post:    Post{Title: "", Content: "Content"},
			wantErr: true,
		},
		{
			name:    "empty content",
			post:    Post{Title: "Title", Content: ""},
			wantErr: true,
		},
		{
			name:    "both empty",
			post:    Post{Title: "", Content: ""},
			wantErr: true,
		},
		{
			name:    "valid post",
			post:    Post{Title: "Valid Title", Content: "Valid Content"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := store.Add(tt.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddPostIDGeneration(t *testing.T) {
	store := NewStore()

	post1 := Post{Title: "Post 1", Content: "Content 1"}
	post2 := Post{Title: "Post 2", Content: "Content 2"}

	store.Add(post1)
	store.Add(post2)

	posts := store.GetAll()
	if posts[0].ID == posts[1].ID {
		t.Error("Expected unique IDs for different posts")
	}
	if posts[0].ID <= 0 || posts[1].ID <= 0 {
		t.Error("Expected positive IDs")
	}
}

func TestAddPostAuthor(t *testing.T) {
	store := NewStore()

	// Test with author provided
	postWithAuthor := Post{
		Title:   "Test Post",
		Content: "Test Content",
		Author:  "Test Author",
	}

	err := store.Add(postWithAuthor)
	if err != nil {
		t.Fatalf("Add() failed: %v", err)
	}

	posts := store.GetAll()
	if posts[0].Author != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", posts[0].Author)
	}
}
