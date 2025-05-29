package decompress

import (
	"compress/gzip"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

// Decompresses an input stream with the `compress/gzip` algorithm.
func UnGzip(input io.Reader, outputFile os.FileInfo, progressBar *progressbar.ProgressBar) (int64, error) {
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

	return io.Copy(io.MultiWriter(w, progressBar), r)
}
