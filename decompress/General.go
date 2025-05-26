package decompress

import (
	"errors"
	"io"
	"os"
)

// In this package, files need to be read, processed, and then written. The
// buffer size determines how many bytes are processed at a time.
const DEFAULT_BUFFER_SIZE = 1024

// Decompresses data read from `input` and writes the results to `output` using
// the algorithm specified. The algorithm should be specified as a string that
// names the Go standard library package that applies to it. I.e., the
// accepted values are:
// - "archive/tar"
// - "archive/zip"
// - "compress/bzip2"
// - "compress/flate"
// - "compress/gzip"
// - "compress/lzw"
// - "compress/zlib"
//
// The `filetypes` package contains functions to associate file names and MIME
// types to each of these algorithms.
func General(algorithm string, input io.Reader, outputDir string) (int, error) {
	_, err := os.Stat(outputDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return 0, err
		}
	}
	output, err := os.Stat(outputDir)
	if err != nil {
		return 0, err
	}

	switch algorithm {
	case "archive/tar":
		return UnTar(input, output)
	default:
		return 0, errors.New("unrecognized compression algorithm\"" + algorithm + "\"")
	}
}