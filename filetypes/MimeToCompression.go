package filetypes

var MimeToCompression = map[string]string{
	"application/x-abiword": "compress/gzip",
	"application/vnd.android.package-archive": "archive/zip",
	"application/x-bzip2": "compress/bzip2",
	"application/x-bzip": "compress/bzip2",
	"application/epub+zip": "archive/zip",
	"image/gif": "compress/lzw",
	"application/x-gzip": "archive/tar + compress/gzip",
	"application/gzip": "archive/tar + compress/gzip",
	"application/java-archive": "archive/zip",
	"application/vnd.google-earth.kmz": "archive/zip",
	"text/x-python": "archive/zip",
	"image/svg+xml": "compress/flate",
	"application/x-tar": "archive/tar",
	"application/x-ms-wmz": "archive/zip",
	"application/x-xpinstall": "archive/zip",
	"application/zip": "archive/zip",
	"application/x-zip-compressed": "archive/zip",
	"application/zip-compressed": "archive/zip",
	"": "undefined",
}
