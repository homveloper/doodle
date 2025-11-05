# Blog Doodle - Real-time Search with Templ & HTMX

A modern blog application demonstrating real-time search functionality using Go, Templ templates, and HTMX.

## Features

- **Real-time Search**: Search posts as you type with instant results
- **HTMX Integration**: Dynamic content updates without page reloads
- **Type-safe Templates**: Templ provides compile-time safety for HTML generation
- **Responsive Design**: Clean, modern UI that works on all devices
- **Zero JavaScript**: All interactivity powered by HTMX attributes

## Tech Stack

- **Go 1.21+**: Backend server and business logic
- **Templ**: Type-safe HTML templating
- **HTMX 1.9**: Dynamic HTML over the wire
- **Standard Library**: No external web framework required

## Architecture

```
blog-templ/
├── models/          # Data models and business logic
│   ├── post.go      # Post struct and Store
│   └── post_test.go # Model tests
├── handlers/        # HTTP handlers
│   ├── handlers.go      # Request handlers
│   └── handlers_test.go # Handler tests
├── templates/       # Templ templates
│   ├── layout.templ # Base layout with styles
│   ├── index.templ  # Home page with search
│   └── posts.templ  # Post list and cards
├── main.go          # Application entry point
└── go.mod           # Go module definition
```

## How It Works

### Real-time Search Flow

1. User types in the search input
2. HTMX intercepts keyup events (with 300ms debounce)
3. GET request sent to `/search?q=<query>`
4. Server filters posts and renders `PostList` template
5. HTMX swaps the `#post-list` content with new results
6. Smooth transition with loading indicator

### Search Implementation

The search functionality matches across multiple fields:
- Post title
- Post content
- Author name
- Tags

All searches are case-insensitive for better user experience.

### HTMX Attributes Used

```html
<input
    hx-get="/search"              <!-- Endpoint to call -->
    hx-trigger="keyup changed delay:300ms"  <!-- Trigger with debounce -->
    hx-target="#post-list"        <!-- Element to update -->
    hx-indicator="#search-indicator"  <!-- Loading indicator -->
/>
```

## Installation

### Prerequisites

- Go 1.21 or higher
- Templ CLI (for template generation)

### Install Templ

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Setup

```bash
# Navigate to the feature directory
cd features/blog-templ

# Download dependencies
go mod download

# Generate Templ templates
templ generate
```

## Running the Application

### Development Mode

```bash
# Generate templates and run
templ generate && go run main.go
```

The server will start on `http://localhost:8080`

### Production Build

```bash
# Generate templates
templ generate

# Build binary
go build -o blog-server

# Run
./blog-server
```

## Running Tests

### All Tests

```bash
go test ./...
```

### With Coverage

```bash
go test -cover ./...
```

### Verbose Output

```bash
go test -v ./...
```

### Specific Package

```bash
# Test models only
go test ./models

# Test handlers only
go test ./handlers
```

## Usage Examples

### Search by Title
Type "templ" to find posts about Templ

### Search by Tag
Type "htmx" to find posts tagged with HTMX

### Search by Author
Type "jane" to find posts by Jane Doe

### Search by Content
Type "web development" to find posts discussing web development

## Sample Data

The application comes with 4 sample blog posts covering:
1. Getting Started with Templ and HTMX
2. Building Real-time Search with HTMX
3. Why Go is Great for Web Development
4. Type-Safe HTML Templates

You can modify the sample data in `models/post.go` in the `NewStore()` function.

## Customization

### Adding More Posts

Edit `models/post.go` and add posts to the `NewStore()` function:

```go
{
    ID:        5,
    Title:     "Your Post Title",
    Content:   "Post content here...",
    Author:    "Author Name",
    CreatedAt: time.Now(),
    Tags:      []string{"tag1", "tag2"},
}
```

### Styling

Styles are embedded in the Templ templates:
- `templates/layout.templ` - Global styles and layout
- `templates/posts.templ` - Post card styles

### Changing Port

Edit `main.go`:

```go
port := 3000  // Change from 8080
```

## Key Features Demonstrated

### 1. Type-safe Templates
Templ generates Go code from templates, providing:
- Compile-time error checking
- IDE autocomplete support
- Type-safe data passing

### 2. HTMX Patterns
- Debounced input handling
- Partial page updates
- Loading indicators
- Progressive enhancement

### 3. Clean Architecture
- Separation of concerns (models, handlers, templates)
- Testable components
- No framework dependencies

### 4. Performance
- No JavaScript bundle to download
- Server-rendered HTML
- Minimal network overhead
- Fast response times

## Testing Strategy

### Model Tests (`models/post_test.go`)
- Store initialization
- Search functionality across all fields
- Case-insensitive matching
- Edge cases (empty queries, no results)

### Handler Tests (`handlers/handlers_test.go`)
- HTTP endpoint responses
- Search query handling
- HTML content verification
- Special character handling

## Performance Considerations

- **Debouncing**: 300ms delay prevents excessive requests
- **In-memory Storage**: Fast search without database overhead
- **Partial Updates**: Only search results are re-rendered
- **No JavaScript Framework**: Minimal client-side overhead

## Future Enhancements

Potential improvements for this doodle:
- Pagination for large result sets
- Advanced filtering (by date, author, tags)
- Persistent storage (database integration)
- Post detail pages
- Markdown support for content
- User authentication
- CRUD operations for posts

## Learning Resources

- [Templ Documentation](https://templ.guide/)
- [HTMX Documentation](https://htmx.org/)
- [Go Web Programming](https://golang.org/doc/)

## License

Part of the Doodle project - experimental features for learning and exploration.
