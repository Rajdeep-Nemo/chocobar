package sleepycat

import (
	"strings"
)

// String reads a string from the input source.
// MaxAttempt does not affect String because every input is valid.
func String(opts ...option) (string, error) {
	return input(func(text string) (string, error) {
		return strings.TrimSpace(text), nil
	}, opts...)
}
