package filetypes

import (
	"strings"
)

// Returns an appropriate MIME type based on the extension on the provided
// filename. If one cannot be found, then it returns an empty string.
func MimeFromName(filename string) string {
	filenameParts := strings.Split(filename, ".")
	fileExtension := filenameParts[len(filenameParts) - 1]

	mime, ok := ExtToMime["." + fileExtension]
	if !ok {
		return ""
	}
	return mime
}
