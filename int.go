package sleepycat

import (
	"strconv"
)

// Reads an integer from the input source and returns as `int`
func Int(opts ...option) (int, error) {
	return input(strconv.Atoi, opts...)
}
