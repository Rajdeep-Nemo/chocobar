package chocobar

import (
	"strconv"
)

// Parses a string and returns a uint32.
func parseUint32(s string) (uint32, error) {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

// Uint32 reads an unsigned integer from the input source and returns it as a uint32.
func Uint32(opts ...option) (uint32, error) {
	return input(parseUint32, opts...)
}
