package decompress

import (
	"io"
)

// Writes the contents of `in` to `out` until `in` stops providing bytes.
// Returns the number of bytes written.
func writeAll(in io.Reader, out io.Writer) (int, error) {
	buffer := make([]byte, DEFAULT_BUFFER_SIZE)

	totalBytesWritten := 0

	for {
		bytesRead, err := in.Read(buffer)
		if err != nil && err != io.EOF {
			return totalBytesWritten, err
		}
		if bytesRead == 0 {
			return totalBytesWritten, nil
		}

		bytesWritten, err := out.Write(buffer[:bytesRead])
		totalBytesWritten += bytesWritten
		if err != nil {
			return totalBytesWritten, err
		}
	}
}
