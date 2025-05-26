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
				return totalBytesWritten, nil
			}
			if bytesRead == 0 {
				break
			}
			
			bytesWritten, err := w.Write(buffer[:bytesRead])
			totalBytesWritten += bytesWritten
			if err != nil {
				return totalBytesWritten, err
			}
		}

		w.Close()
	}
}
