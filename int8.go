package chocobar

import (
	"strconv"
)

// Parses a string and returns 8 bit integer
func parseInt8(s string) (int8, error) {
	val, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(val), nil
}

// Int8 reads an integer from the input source and returns it as an int8.
func Int8(opts ...option) (int8, error) {
	return input(parseInt8, opts...)
}
