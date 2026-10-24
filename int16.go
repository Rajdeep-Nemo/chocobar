package chocobar

import (
	"strconv"
)

// Parses a string and returns 16 bit integer
func parseInt16(s string) (int16, error) {
	val, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(val), nil
}

// Int16 reads an integer from the input source and returns it as an int16.
func Int16(opts ...option) (int16, error) {
	return input(parseInt16, opts...)
}
