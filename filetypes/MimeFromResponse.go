package filetypes

import (
	"net/http"
	"strings"
)

// Returns the MIME type from an HTTP response. If a type cannot be determined,
// then this will return an empty string.
func MimeFromResponse(resp *http.Response) string {
	contentTypeHeader := resp.Header.Get("Content-Type")

	if contentTypeHeader == "" {
		return ""
	}

	// Some `Content-Type` headers will be in a form like 
	// "text/html; Encoding=...". However, we only want the "text/html" part.
	contentTypeHeaderParts := strings.Split(contentTypeHeader, ";")
	return contentTypeHeaderParts[0]
}
