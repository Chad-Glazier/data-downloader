package decompress

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Chad-Glazier/fdd/filetypes"
)

// Decompresses a file, writing it to the named destination. If the the
// destination already exists, then it will be used. The number of bytes in the
// resulting (decompressed) file is returned.
//
// The `algorithm` argument should be the name of a Go standard library package
// that should be used. I.e., the acceptable arguments are:
// - "archive/tar"
// - "archive/zip"
// - "compress/bzip2"
// - "compress/flate"
// - "compress/gzip"
// - "compress/zlib"
//
// Note that in order to decompress a single file (e.g., using the Gzip
// algorithm), the destination must be a file. In order to decompress an
// archive (e.g., a .tar or .zip file), the destination must be a directory.
func File(source string, destination string) (int, error) {
	inputFile, err := os.Open(source)
	if err != nil {
		return 0, nil
	}
	defer inputFile.Close()

	//
	// Try using the file extension to get an algorithm.
	//

	ext := filepath.Ext(source)
	if ext != "" {
		fileType := filetypes.ExtToMime[ext]
		algorithm := filetypes.MimeToCompression[fileType]
		if algorithm != "" {
			return General(algorithm, inputFile, destination)
		}
	}

	//
	// Try using the magic number of the file to get an algorithm.
	//

	headerReader, err := os.Open(source)
	if err != nil {
		return 0, nil
	}

	algorithm, _ := filetypes.CompressionFromMagicNum(headerReader)
	headerReader.Close()
	if algorithm != "" {
		return General(algorithm, inputFile, destination)
	}

	//
	// The previous methods failed; return an error.
	//

	return 0, fmt.Errorf("could not determine file type of %s", source)
}
