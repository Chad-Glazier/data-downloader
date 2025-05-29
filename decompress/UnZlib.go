package decompress

import (
	"compress/zlib"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

// Decompresses a file with the `compress/zlib` algorithm.
func UnZlib(input io.Reader, outputFile os.FileInfo, progressBar *progressbar.ProgressBar) (int64, error) {
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

	return io.Copy(io.MultiWriter(w, progressBar), r)
}
