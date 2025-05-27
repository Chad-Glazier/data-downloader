package misc

import (
	"strings"
	"path/filepath"
)

// Returns the base name of a file. Specifically, it returns the file path's
// base name with the last extension removed. E.g.,
// - "path/to/file.ext" becomes "file"
// - "path/to/file.tar.gz" becomes "file.tar"
func BaseName(filename string) string {
	// If `filename` is a path string (i.e., one with slashes) then only take
	// the last part. Note that this function preserves the extension.
	filename = filepath.Base(filename)

	// Remove the last extension, if one exists.
	filenameParts := strings.Split(filename, ".")
	switch len(filenameParts) {
	case 1, 2:
		return filenameParts[0]
	default:
		return strings.Join(filenameParts[:len(filenameParts) - 2], ".")
	}
}
