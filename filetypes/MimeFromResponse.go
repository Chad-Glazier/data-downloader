package filetypes

import (
	"net/http"
	"strings"

	"github.com/Chad-Glazier/fdd/misc"
)

// Returns the MIME type from an HTTP response. If a type cannot be determined,
// then this will return an empty string.
//
// This function attempts to, first, determine the filetype from the request
// URL's extension. E.g., a request sent to "https://data.com/sample.json" will
// be presumed to respond with a file of the type "application/json". If such
// an extension doesn't exist (or is not recognized), then this function defers
// to the "Content-Type" header of the response.
func MimeFromResponse(resp *http.Response) string {
	filename, _ := misc.FileNameFromUrl(resp.Request.URL.String())
	filenameParts := strings.Split(filename, ".")
	if len(filenameParts) != 1 {
		extension := "." + filenameParts[len(filenameParts)-1]
		return ExtToMime[extension]
	}

	contentTypeHeader := resp.Header.Get("Content-Type")
	// Some `Content-Type` headers will be in a form like
	// "text/html; Encoding=...". However, we only want the "text/html" part.
	contentTypeHeaderParts := strings.Split(contentTypeHeader, ";")
	return contentTypeHeaderParts[0]
}
