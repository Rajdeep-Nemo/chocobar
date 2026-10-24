package chocobar

import (
	"fmt"
)

// Parses the string value
func parseByte(s string) (byte, error) {
	if len(s) != 1 {
		return 0, fmt.Errorf("chocobar: parsing %q: expected a single ASCII character", s)
	}
	return s[0], nil
}

// Byte reads a single ASCII character from the input source and returns it as a byte.
func Byte(opts ...option) (byte, error) {
	return input(parseByte, opts...)
}
