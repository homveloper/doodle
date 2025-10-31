package asciiart

import (
	"strings"
	"testing"
)

func TestGenerate_Basic(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(string) bool
	}{
		{
			name:  "single letter",
			input: "A",
			check: func(s string) bool {
				return len(s) > 0 && strings.Contains(s, "#")
			},
		},
		{
			name:  "multiple letters",
			input: "ABC",
			check: func(s string) bool {
				lines := strings.Split(s, "\n")
				return len(lines) == 5 // Standard font height
			},
		},
		{
			name:  "numbers",
			input: "123",
			check: func(s string) bool {
				return strings.Contains(s, "#")
			},
		},
		{
			name:  "mixed alphanumeric",
			input: "A1B2",
			check: func(s string) bool {
				lines := strings.Split(s, "\n")
				return len(lines) == 5
			},
		},
		{
			name:  "empty string",
			input: "",
			check: func(s string) bool {
				return s == ""
			},
		},
		{
			name:  "space",
			input: " ",
			check: func(s string) bool {
				lines := strings.Split(s, "\n")
				return len(lines) == 5
			},
		},
		{
			name:  "word with space",
			input: "HI THERE",
			check: func(s string) bool {
				lines := strings.Split(s, "\n")
				return len(lines) == 5 && len(lines[0]) > 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Generate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.check(result) {
				t.Errorf("Generate() validation failed for input %q\nGot:\n%s", tt.input, result)
			}
		})
	}
}

func TestGenerate_WithBorder(t *testing.T) {
	result, err := Generate("GO", WithBorder())
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	lines := strings.Split(result, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected at least 3 lines with border, got %d", len(lines))
	}

	// Check for border characters
	if !strings.HasPrefix(lines[0], "╔") {
		t.Errorf("Expected top border to start with ╔, got: %s", lines[0])
	}
	if !strings.HasPrefix(lines[len(lines)-1], "╚") {
		t.Errorf("Expected bottom border to start with ╚, got: %s", lines[len(lines)-1])
	}
}

func TestGenerate_WithPadding(t *testing.T) {
	withoutPadding, _ := Generate("A")
	withPadding, _ := Generate("A", WithPadding(5))

	linesWithout := strings.Split(withoutPadding, "\n")
	linesWith := strings.Split(withPadding, "\n")

	if len(linesWithout) != len(linesWith) {
		t.Errorf("Line count mismatch")
	}

	// Check that padded version is wider
	if len(linesWith[0]) <= len(linesWithout[0]) {
		t.Errorf("Padded version should be wider")
	}
}

func TestGenerate_WithAlignment(t *testing.T) {
	tests := []struct {
		name  string
		align Align
		width int
	}{
		{"left", AlignLeft, 40},
		{"center", AlignCenter, 40},
		{"right", AlignRight, 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Generate("A", WithAlignment(tt.align), WithWidth(tt.width))
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			lines := strings.Split(result, "\n")
			if len(lines) == 0 {
				t.Fatal("No lines generated")
			}

			// Just verify it doesn't error and produces output
			if len(result) == 0 {
				t.Error("Expected non-empty result")
			}
		})
	}
}

func TestGenerate_WithStyle(t *testing.T) {
	tests := []struct {
		name  string
		style Style
	}{
		{"normal", StyleNormal},
		{"shadow", StyleShadow},
		{"double", StyleDouble},
		{"dotted", StyleDotted},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Generate("A", WithStyle(tt.style))
			if err != nil {
				t.Fatalf("Generate() error = %v", err)
			}

			if len(result) == 0 {
				t.Error("Expected non-empty result")
			}
		})
	}
}

func TestGenerate_CombinedOptions(t *testing.T) {
	result, err := Generate("HELLO",
		WithFont(FontStandard),
		WithBorder(),
		WithPadding(2),
		WithAlignment(AlignCenter),
		WithWidth(60),
	)

	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	lines := strings.Split(result, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected multiple lines, got %d", len(lines))
	}

	// Verify border exists
	if !strings.HasPrefix(lines[0], "╔") {
		t.Error("Expected border at top")
	}
}

func TestGenerate_UnsupportedCharacters(t *testing.T) {
	// Should not error on unsupported characters, should use placeholder
	result, err := Generate("A§B", WithFont(FontStandard))
	if err != nil {
		t.Errorf("Unexpected error for unsupported character: %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected non-empty result even with unsupported characters")
	}
}

func TestGenerate_AllLetters(t *testing.T) {
	// Test that all uppercase letters work
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result, err := Generate(alphabet)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected non-empty result for alphabet")
	}

	lines := strings.Split(result, "\n")
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines for standard font, got %d", len(lines))
	}
}

func TestGenerate_AllNumbers(t *testing.T) {
	// Test that all numbers work
	numbers := "0123456789"
	result, err := Generate(numbers)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected non-empty result for numbers")
	}

	lines := strings.Split(result, "\n")
	if len(lines) != 5 {
		t.Errorf("Expected 5 lines for standard font, got %d", len(lines))
	}
}

func TestGenerate_CommonPunctuation(t *testing.T) {
	// Test common punctuation marks
	punct := "!?.,;:'-"
	result, err := Generate(punct)
	if err != nil {
		t.Fatalf("Generate() error = %v", err)
	}

	if len(result) == 0 {
		t.Error("Expected non-empty result for punctuation")
	}
}

// Benchmark tests
func BenchmarkGenerate_Short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Generate("HELLO")
	}
}

func BenchmarkGenerate_Long(b *testing.B) {
	text := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
	for i := 0; i < b.N; i++ {
		_, _ = Generate(text)
	}
}

func BenchmarkGenerate_WithOptions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Generate("TEST",
			WithBorder(),
			WithPadding(2),
			WithAlignment(AlignCenter),
			WithWidth(40),
		)
	}
}
