package chocobar

import (
	"strconv"
)

// Parses a string and returns a uint8.
func parseUint8(s string) (uint8, error) {
	val, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(val), nil
}

// Uint8 reads an unsigned integer from the input source and returns it as a uint8.
func Uint8(opts ...option) (uint8, error) {
	return input(parseUint8, opts...)
}
