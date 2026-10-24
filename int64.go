package chocobar

import (
	"strconv"
)

// Parses a string and returns 64 bit integer
func parseInt64(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// Int64 reads an integer from the input source and returns it as an int64.
func Int64(opts ...option) (int64, error) {
	return input(parseInt64, opts...)
}
