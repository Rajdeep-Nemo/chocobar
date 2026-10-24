package chocobar

import (
	"strconv"
)

// Parses a string and returns a uint64.
func parseUint64(s string) (uint64, error) {
	val, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

// Uint64 reads an unsigned integer from the input source and returns it as a uint64.
func Uint64(opts ...option) (uint64, error) {
	return input(parseUint64, opts...)
}
