package sleepycat

import (
	"strconv"
)

// Parses the string value and returns a uint
func parseUint(s string) (uint, error) {
	val, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}

// Uint reads an unsigned integer from the input source and returns it as a uint.
func Uint(opts ...option) (uint, error) {
	return input(parseUint, opts...)
}
