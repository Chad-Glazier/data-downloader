package decompress

import (
	"io"
	"os"
	"compress/gzip"
)

// Decompresses an input stream with the `compress/gzip` algorithm.
func UnGzip(input io.Reader, outputFile os.FileInfo) (int, error) {
	r, err := gzip.NewReader(input)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	w, err := os.Create(outputFile.Name())
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return writeAll(r, w)
}
