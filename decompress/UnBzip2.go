package decompress

import (
	"compress/bzip2"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

// Decompresses a file with the `compress/bzip` algorithm.
func UnBzip2(input io.Reader, outputFile os.FileInfo, progressBar *progressbar.ProgressBar) (int64, error) {
	r := bzip2.NewReader(input)

	w, err := os.Create(outputFile.Name())
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return io.Copy(io.MultiWriter(w, progressBar), r)
}
