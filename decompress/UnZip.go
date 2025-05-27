package decompress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Decompresses a file with the `archive/zip` algorithm and writes it to the
// given directory.
func UnZip(input io.Reader, outputDir os.FileInfo) (int, error) {
	// .zip files require random access, so there's no way to simply pipe data
	// from the input to the decompression algorithm. Instead, we need to
	// create a temporary file.
	cwd, err := os.Getwd()
	if err != nil {
		return 0, err
	}

	//
	// Downloading the `.zip` archive
	//

	compressed, err := os.CreateTemp(cwd, "data-downloader_*.zip")
	if err != nil {
		return 0, err
	}
	defer os.Remove(compressed.Name())
	defer compressed.Close()

	compressedSize, err := writeAll(input, compressed)
	if err != nil {
		return 0, err
	}

	//
	// Decompressing the `.zip` archive.
	//

	archive, err := zip.NewReader(compressed, int64(compressedSize))
	if err != nil {
		return 0, nil
	}

	totalBytesWritten := 0
	for _, file := range archive.File {

		//
		// Handle the case where it's a directory.
		//

		if file.Mode().IsDir() {
			// Check if the directory already exists. If it doesn't, then
			// create it.
			_, err := os.Stat(filepath.Join(
				outputDir.Name(), file.Name,
			))
			if !os.IsNotExist(err) {
				continue
			}
			err = os.Mkdir(filepath.Join(outputDir.Name(), file.Name), 0755)
			if err != nil {
				return totalBytesWritten, err
			}
			continue
		}

		//
		// Handle the case where it's a file.
		//

		w, err := os.Create(filepath.Join(outputDir.Name(), file.Name))
		if err != nil {
			return totalBytesWritten, err
		}
		r, err := file.Open()
		if err != nil {
			return totalBytesWritten, err
		}

		bytesWritten, err := writeAll(r, w)
		w.Close()
		r.Close()
		totalBytesWritten += bytesWritten
		if err != nil {
			return totalBytesWritten, err
		}
	}

	return totalBytesWritten, nil
}
