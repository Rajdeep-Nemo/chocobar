package sleepycat

import (
	"strconv"
)

// Parses a string and returns a uint16.
func parseUint16(s string) (uint16, error) {
	val, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(val), nil
}

// Uint16 reads an unsigned integer from the input source and returns it as a uint16.
func Uint16(opts ...option) (uint16, error) {
	return input(parseUint16, opts...)
}
