package chocobar

import (
	"strconv"
)

// Int reads an integer from the input source and returns it as an int.
func Int(opts ...option) (int, error) {
	return input(strconv.Atoi, opts...)
}
