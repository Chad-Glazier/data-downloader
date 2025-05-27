package misc

import (
	"strings"
)

// Returns `true` if the filename given has an extension, and `false` otherwise.
func HasExtension(filename string) bool {
	filenameParts := strings.Split(filename, ".")
	return len(filenameParts) > 1
}