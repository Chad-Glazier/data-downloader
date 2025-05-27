package decompress

import (
	"compress/bzip2"
	"io"
	"os"
)

// Decompresses a file with the `compress/bzip` algorithm.
func UnBzip2(input io.Reader, outputFile os.FileInfo) (int, error) {
	r := bzip2.NewReader(input)

	w, err := os.Create(outputFile.Name())
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return writeAll(r, w)
}
