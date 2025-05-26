package misc

import (
	"strings"
)

// Returns the base name of a file.
func BaseName(filename string) string {
	filenameParts := strings.Split(filename, ".")
	return filenameParts[0]
}
