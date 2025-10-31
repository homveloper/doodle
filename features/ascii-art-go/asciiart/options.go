package asciiart

// Font represents the ASCII art font style
type Font string

const (
	// FontStandard is the default 5-line height font
	FontStandard Font = "standard"
	// FontBig is a large, bold font style
	FontBig Font = "big"
	// FontSmall is a compact 3-line height font
	FontSmall Font = "small"
	// FontBlock is a block-style font
	FontBlock Font = "block"
	// FontBanner is a banner-style font
	FontBanner Font = "banner"
)

// Style represents the visual style of the ASCII art
type Style string

const (
	// StyleNormal is the default style with no effects
	StyleNormal Style = "normal"
	// StyleShadow adds shadow effects
	StyleShadow Style = "shadow"
	// StyleDouble uses double-line characters
	StyleDouble Style = "double"
	// StyleDotted uses dotted line style
	StyleDotted Style = "dotted"
)

// Align represents text alignment
type Align string

const (
	// AlignLeft aligns text to the left (default)
	AlignLeft Align = "left"
	// AlignCenter centers the text
	AlignCenter Align = "center"
	// AlignRight aligns text to the right
	AlignRight Align = "right"
)

// Config holds the configuration for ASCII art generation
type Config struct {
	Font      Font  // Font style to use
	Width     int   // Maximum width (0 = unlimited)
	Padding   int   // Left and right padding
	Alignment Align // Text alignment
	Border    bool  // Whether to add a border
	Style     Style // Visual style
}

// defaultConfig returns a Config with default values
func defaultConfig() *Config {
	return &Config{
		Font:      FontStandard,
		Width:     0,
		Padding:   0,
		Alignment: AlignLeft,
		Border:    false,
		Style:     StyleNormal,
	}
}

// Option is a function that modifies a Config
type Option func(*Config)

// WithFont sets the font style
func WithFont(font Font) Option {
	return func(c *Config) {
		c.Font = font
	}
}

// WithWidth sets the maximum width
func WithWidth(width int) Option {
	return func(c *Config) {
		if width > 0 {
			c.Width = width
		}
	}
}

// WithPadding sets the left and right padding
func WithPadding(padding int) Option {
	return func(c *Config) {
		if padding >= 0 {
			c.Padding = padding
		}
	}
}

// WithAlignment sets the text alignment
func WithAlignment(align Align) Option {
	return func(c *Config) {
		c.Alignment = align
	}
}

// WithBorder enables border around the ASCII art
func WithBorder() Option {
	return func(c *Config) {
		c.Border = true
	}
}

// WithStyle sets the visual style
func WithStyle(style Style) Option {
	return func(c *Config) {
		c.Style = style
	}
}
