package chocobar

import (
	"strconv"
)

// Bool reads a boolean value from the input source.
func Bool(opts ...option) (bool, error) {
	return input(strconv.ParseBool, opts...)
}
