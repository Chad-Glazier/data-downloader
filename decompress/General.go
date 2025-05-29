package decompress

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
)

// Decompresses data read from `input` and writes the results to `output` using
// the algorithm specified. The algorithm should be specified as a string that
// names the Go standard library package that applies to it. I.e., the
// accepted values are:
// - "archive/tar"
// - "archive/zip"
// - "compress/bzip2"
// - "compress/flate"
// - "compress/gzip"
// - "compress/zlib"
//
// (Note that "compress/lzw" is omitted from the list.)
//
// `destination` is created in the current working directory. If the file
// already exists, it will be used. However, if you're decompressing an archive
// the existing file must be a directory. Conversely, if you're decompressing
// with a single-file algorithm (the "compress/*" algorithms), then the existing
// destination must be a non-directory. In the case of single-file decompression,
// the existing file is overwritten. In the case of an archive, the existing
// directory will not be cleared.
//
// The `progressBar` argument is the progress bar that will be used to display
// the progress (duh). If you don't want a progress bar, just pass `nil`.
//
func General(algorithm string, input io.Reader, destination string, progressBar *progressbar.ProgressBar) (int64, error) {
	algorithmParts := strings.Split(algorithm, "/")
	classification := algorithmParts[0]
	algorithm = algorithmParts[len(algorithmParts) - 1]

	var absoluteDest string
	cwd, err := os.Getwd()
	if err != nil {
		absoluteDest = destination
	} else {
		absoluteDest = filepath.Join(cwd, destination)
	}

	var output os.FileInfo

	switch classification {
	case "archive": 
		// We're decompressing an archive (i.e., many files). The output
		// should be a new directory.
		_, err := os.Stat(destination)
		if os.IsNotExist(err) {
			err := os.MkdirAll(destination, 0755)
			if err != nil {
				return 0, err
			}
		}
		output, err = os.Stat(destination)
		if err != nil {
			return 0, err
		}
		if !output.IsDir() {
			return 0, fmt.Errorf(
				"the destination file %s already exists, but it is not a directory",
				absoluteDest,
			)
		}
	case "compress":
		// We're decompressing a single file. The output should be a new
		// file.
		_, err := os.Stat(destination)
		if os.IsNotExist(err) {
			destFile, err := os.Create(destination)
			if err != nil {
				return 0, err
			}
			destFile.Close()
		}
		output, err = os.Stat(destination)
		if err != nil {
			return 0, err
		}
		if output.IsDir() {
			return 0, fmt.Errorf(
				"the destination %s already exists, but it is a directory",
				absoluteDest,
			)
		}
	default:
		return 0, errors.New("unsupported compression algorithm\"" + algorithm + "\"")
	}

	switch algorithm {
	case "tar":
		return UnTar(input, output, progressBar)
	case "zip":
		return UnZip(input, output, progressBar)
	case "bzip2":
		return UnBzip2(input, output, progressBar)
	case "flate":
		return UnFlate(input, output, progressBar)
	case "gzip":
		return UnGzip(input, output, progressBar)
	case "zlib":
		return UnZlib(input, output, progressBar)
	default:
		return 0, errors.New("unsupported compression algorithm\"" + algorithm + "\"")
	}
}
