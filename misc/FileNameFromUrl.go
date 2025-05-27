package misc

import (
	"strings"
	"net/url"
)

// Returns the filename from a URL, if one exists. E.g.,
// - "https://google.com/file.ext?q=a" returns "file.ext"
// - "https://google.com/file?q" returns "file"
// - "https://google.com" returns ""
// 
// Returns an error if the URL is invalid.
func FileNameFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	
	pathParts := strings.Split(u.Path, "/")
	return pathParts[len(pathParts) - 1], nil
}
