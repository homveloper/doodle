package main

import (
	"fmt"

	"github.com/homveloper/doodle/features/ascii-art-go/asciiart"
)

func main() {
	fmt.Println("=== ASCII Art Generator Examples ===\n")

	// Example 1: Basic usage
	fmt.Println("1. Basic Usage:")
	fmt.Println("---")
	result, _ := asciiart.Generate("HELLO")
	fmt.Println(result)
	fmt.Println()

	// Example 2: With border
	fmt.Println("2. With Border:")
	fmt.Println("---")
	result, _ = asciiart.Generate("GO", asciiart.WithBorder())
	fmt.Println(result)
	fmt.Println()

	// Example 3: With padding and border
	fmt.Println("3. With Padding and Border:")
	fmt.Println("---")
	result, _ = asciiart.Generate("HI",
		asciiart.WithPadding(3),
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 4: Center alignment
	fmt.Println("4. Center Alignment:")
	fmt.Println("---")
	result, _ = asciiart.Generate("CENTER",
		asciiart.WithAlignment(asciiart.AlignCenter),
		asciiart.WithWidth(60),
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 5: Right alignment
	fmt.Println("5. Right Alignment:")
	fmt.Println("---")
	result, _ = asciiart.Generate("RIGHT",
		asciiart.WithAlignment(asciiart.AlignRight),
		asciiart.WithWidth(60),
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 6: Shadow style
	fmt.Println("6. Shadow Style:")
	fmt.Println("---")
	result, _ = asciiart.Generate("SHADOW",
		asciiart.WithStyle(asciiart.StyleShadow),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 7: Double style
	fmt.Println("7. Double Style:")
	fmt.Println("---")
	result, _ = asciiart.Generate("DOUBLE",
		asciiart.WithStyle(asciiart.StyleDouble),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 8: Dotted style
	fmt.Println("8. Dotted Style:")
	fmt.Println("---")
	result, _ = asciiart.Generate("DOTTED",
		asciiart.WithStyle(asciiart.StyleDotted),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 9: Numbers
	fmt.Println("9. Numbers:")
	fmt.Println("---")
	result, _ = asciiart.Generate("2025", asciiart.WithBorder())
	fmt.Println(result)
	fmt.Println()

	// Example 10: Mixed content
	fmt.Println("10. Mixed Content:")
	fmt.Println("---")
	result, _ = asciiart.Generate("GO 1.23",
		asciiart.WithAlignment(asciiart.AlignCenter),
		asciiart.WithWidth(50),
		asciiart.WithPadding(2),
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 11: Punctuation
	fmt.Println("11. With Punctuation:")
	fmt.Println("---")
	result, _ = asciiart.Generate("HELLO!",
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	// Example 12: Full alphabet
	fmt.Println("12. Full Alphabet (A-M):")
	fmt.Println("---")
	result, _ = asciiart.Generate("ABCDEFGHIJKLM")
	fmt.Println(result)
	fmt.Println()

	fmt.Println("13. Full Alphabet (N-Z):")
	fmt.Println("---")
	result, _ = asciiart.Generate("NOPQRSTUVWXYZ")
	fmt.Println(result)
	fmt.Println()

	// Example 14: All digits
	fmt.Println("14. All Digits:")
	fmt.Println("---")
	result, _ = asciiart.Generate("0123456789")
	fmt.Println(result)
	fmt.Println()

	// Example 15: Complex combination
	fmt.Println("15. Complex Example:")
	fmt.Println("---")
	result, _ = asciiart.Generate("ASCII ART",
		asciiart.WithStyle(asciiart.StyleDouble),
		asciiart.WithAlignment(asciiart.AlignCenter),
		asciiart.WithWidth(70),
		asciiart.WithPadding(2),
		asciiart.WithBorder(),
	)
	fmt.Println(result)
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
