package chocobar

import (
	"strconv"
)

// Float64 reads a floating-point number from the input source and returns it as a float64.
func Float64(opts ...option) (float64, error) {
	return input(func(s string) (float64, error) {
		return strconv.ParseFloat(s, 64)
	}, opts...)
}
