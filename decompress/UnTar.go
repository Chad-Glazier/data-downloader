package decompress

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// Decompresses a file with the `archive/tar` algorithm and writes it to the
// given directory.
func UnTar(input io.Reader, outputDir os.FileInfo) (int, error) {
	r := tar.NewReader(input)

	buffer := make([]byte, DEFAULT_BUFFER_SIZE)

	totalBytesWritten := 0
	
	for {
		header, err := r.Next()
		if err == io.EOF {
			return totalBytesWritten, nil
		}
		if err != nil {
			return totalBytesWritten, err
		}

		if header.Typeflag == tar.TypeDir {
			// If the file is a directory, check if the directory already
			// exists. If it does not, then create it.
			_, err := os.Stat(filepath.Join(
				outputDir.Name(), header.Name,
			))
			if !os.IsNotExist(err) {
				continue
			}
			err = os.Mkdir(filepath.Join(outputDir.Name(), header.Name), 0755)
			if err != nil {
				return totalBytesWritten, err
			}
			continue
		}

		w, err := os.Create(filepath.Join(outputDir.Name(), header.Name))
		if err != nil {
			return totalBytesWritten, err
		}	

		for {
			bytesRead, err := r.Read(buffer)
			if err != nil && err != io.EOF {
				w.Close()
				return totalBytesWritten, err
			}
			if bytesRead == 0 {
				break
			}
			
			bytesWritten, err := w.Write(buffer[:bytesRead])
			totalBytesWritten += bytesWritten
			if err != nil {
				w.Close()
				return totalBytesWritten, err
			}
		}

		w.Close()
	}
}
