package chocobar

import (
	"fmt"
)

// Parses the string value and returns a rune
func parseRune(s string) (rune, error) {
	r := []rune(s)
	if len(r) != 1 {
		return 0, fmt.Errorf("chocobar: parsing %q: expected a single character", s)
	}
	return r[0], nil
}

// Rune reads a Unicode code point from the input source.
func Rune(opts ...option) (rune, error) {
	return input(parseRune, opts...)
}
