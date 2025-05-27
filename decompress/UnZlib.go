package decompress

import (
	"io"
	"os"
	"compress/zlib"
)

// Decompresses a file with the `compress/zlib` algorithm.
func UnZlib(input io.Reader, outputFile os.FileInfo) (int, error) {
	r, err := zlib.NewReader(input)
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
