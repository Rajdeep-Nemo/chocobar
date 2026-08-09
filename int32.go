package sleepycat

import (
	"strconv"
)

// Parses a string and returns 32 bit integer
func parseInt32(s string) (int32, error) {
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// Int32 reads an integer from the input source and returns it as an int32.
func Int32(opts ...option) (int32, error) {
	return input(parseInt32, opts...)
}
