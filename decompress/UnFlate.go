package decompress

import (
	"compress/flate"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
)

// Decompresses a file with the `compress/flate` algorithm.
//
// Yes, I know "inflate" would be better. But all of the other functions start
// with "un".
func UnFlate(input io.Reader, outputFile os.FileInfo, progressBar *progressbar.ProgressBar) (int64, error) {
	r := flate.NewReader(input)
	defer r.Close()

	w, err := os.Create(outputFile.Name())
	if err != nil {
		return 0, err
	}
	defer w.Close()

	return io.Copy(io.MultiWriter(w, progressBar), r)
}
