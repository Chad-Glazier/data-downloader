package decompress

import (
	"io"
	"os"
	"compress/flate"
)

// Decompresses a file with the `compress/flate` algorithm.
//
// Yes, I know "inflate" would be better. But all of the other functions start
// with "un".
func UnFlate(input io.Reader, outputFile os.FileInfo) (int, error) {
	r := flate.NewReader(input)
	defer r.Close()

	w, err := os.Create(outputFile.Name())
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return writeAll(r, w)
}
