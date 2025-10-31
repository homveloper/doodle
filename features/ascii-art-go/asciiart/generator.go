package asciiart

import (
	"fmt"
	"strings"
)

// Generate converts the given text to ASCII art with the specified options
func Generate(text string, opts ...Option) (string, error) {
	// Apply options to default config
	config := defaultConfig()
	for _, opt := range opts {
		opt(config)
	}

	// Handle empty text
	if text == "" {
		return "", nil
	}

	// Get font data
	font, err := getFont(config.Font)
	if err != nil {
		return "", err
	}

	// Convert text to ASCII art lines
	lines, err := textToLines(text, font)
	if err != nil {
		return "", err
	}

	// Apply padding
	if config.Padding > 0 {
		lines = applyPadding(lines, config.Padding)
	}

	// Apply alignment
	if config.Width > 0 && config.Alignment != AlignLeft {
		lines = applyAlignment(lines, config.Width, config.Alignment)
	}

	// Apply style
	lines = applyStyle(lines, config.Style)

	// Add border
	if config.Border {
		lines = addBorder(lines)
	}

	return strings.Join(lines, "\n"), nil
}

// textToLines converts text to ASCII art lines using the given font
func textToLines(text string, font *fontData) ([]string, error) {
	if len(text) == 0 {
		return []string{}, nil
	}

	// Initialize result lines with the height of the font
	result := make([]string, font.height)

	// Process each character
	for _, char := range text {
		// Get ASCII representation for this character
		charLines, ok := font.chars[char]
		if !ok {
			// Use placeholder for unsupported characters
			charLines, ok = font.chars['?']
			if !ok {
				// If no placeholder, skip this character
				continue
			}
		}

		// Append each line of the character to the corresponding result line
		for i := 0; i < font.height; i++ {
			if i < len(charLines) {
				result[i] += charLines[i]
			}
		}
	}

	return result, nil
}

// applyPadding adds left and right padding to each line
func applyPadding(lines []string, padding int) []string {
	if padding <= 0 {
		return lines
	}

	pad := strings.Repeat(" ", padding)
	result := make([]string, len(lines))

	for i, line := range lines {
		result[i] = pad + line + pad
	}

	return result
}

// applyAlignment aligns lines according to the specified alignment and width
func applyAlignment(lines []string, width int, align Align) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		lineLen := len(line)

		// If line is already wider than or equal to width, no alignment needed
		if lineLen >= width {
			result[i] = line
			continue
		}

		padding := width - lineLen

		switch align {
		case AlignCenter:
			leftPad := padding / 2
			rightPad := padding - leftPad
			result[i] = strings.Repeat(" ", leftPad) + line + strings.Repeat(" ", rightPad)
		case AlignRight:
			result[i] = strings.Repeat(" ", padding) + line
		default: // AlignLeft
			result[i] = line + strings.Repeat(" ", padding)
		}
	}

	return result
}

// applyStyle applies the specified style to the lines
func applyStyle(lines []string, style Style) []string {
	switch style {
	case StyleShadow:
		return applyShadowStyle(lines)
	case StyleDouble:
		return applyDoubleStyle(lines)
	case StyleDotted:
		return applyDottedStyle(lines)
	default: // StyleNormal
		return lines
	}
}

// applyShadowStyle adds a shadow effect
func applyShadowStyle(lines []string) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		// Add shadow character after non-space characters
		shadowLine := ""
		for j, ch := range line {
			shadowLine += string(ch)
			// Add shadow if this is not the last character and next char is space
			if ch != ' ' && j+1 < len(line) && line[j+1] == ' ' {
				shadowLine = shadowLine[:len(shadowLine)-1] + string(ch) + "░"
			}
		}
		result[i] = shadowLine
	}

	// Add a shadow line at the bottom
	if len(lines) > 0 {
		shadowBottom := ""
		for _, ch := range lines[len(lines)-1] {
			if ch != ' ' {
				shadowBottom += "░"
			} else {
				shadowBottom += " "
			}
		}
		result = append(result, shadowBottom)
	}

	return result
}

// applyDoubleStyle replaces characters with double-line equivalents
func applyDoubleStyle(lines []string) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		doubleLine := ""
		for _, ch := range line {
			switch ch {
			case '#':
				doubleLine += "█"
			case '-':
				doubleLine += "═"
			case '|':
				doubleLine += "║"
			default:
				doubleLine += string(ch)
			}
		}
		result[i] = doubleLine
	}

	return result
}

// applyDottedStyle replaces characters with dotted equivalents
func applyDottedStyle(lines []string) []string {
	result := make([]string, len(lines))

	for i, line := range lines {
		dottedLine := ""
		for _, ch := range line {
			switch ch {
			case '#':
				dottedLine += "▒"
			case '-':
				dottedLine += "┈"
			case '|':
				dottedLine += "┊"
			default:
				dottedLine += string(ch)
			}
		}
		result[i] = dottedLine
	}

	return result
}

// addBorder adds a border around the lines
func addBorder(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	// Find the maximum width
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	// Normalize all lines to the same width
	normalizedLines := make([]string, len(lines))
	for i, line := range lines {
		if len(line) < maxWidth {
			normalizedLines[i] = line + strings.Repeat(" ", maxWidth-len(line))
		} else {
			normalizedLines[i] = line
		}
	}

	result := make([]string, 0, len(lines)+2)

	// Top border
	result = append(result, "╔"+strings.Repeat("═", maxWidth)+"╗")

	// Content with side borders
	for _, line := range normalizedLines {
		result = append(result, "║"+line+"║")
	}

	// Bottom border
	result = append(result, "╚"+strings.Repeat("═", maxWidth)+"╝")

	return result
}

// getFont returns the font data for the specified font type
func getFont(font Font) (*fontData, error) {
	switch font {
	case FontStandard:
		return getStandardFont(), nil
	case FontBig:
		return nil, fmt.Errorf("font %q not yet implemented", font)
	case FontSmall:
		return nil, fmt.Errorf("font %q not yet implemented", font)
	case FontBlock:
		return nil, fmt.Errorf("font %q not yet implemented", font)
	case FontBanner:
		return nil, fmt.Errorf("font %q not yet implemented", font)
	default:
		return nil, fmt.Errorf("unknown font: %q", font)
	}
}
