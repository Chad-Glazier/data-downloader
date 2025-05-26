package main

import (
	"io"
	"net/http"
	"os"
)

const DEFAULT_BUFFER_SIZE = 1024

// Used to write an uncompressed response body to a file.
//
// Returns the number of bytes written, unless there was an error (in which
// case the number returned is -1).
func writeBodyToFile(filename string, resp *http.Response) (int, error) {
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	var bufferSize int
	if resp.ContentLength == -1 {
		bufferSize = DEFAULT_BUFFER_SIZE
	} else {
		bufferSize = int(min(DEFAULT_BUFFER_SIZE, resp.ContentLength))
	}
	
	buffer := make([]byte, bufferSize)
	bytesWritten := int64(0)
	for {
		bytesRead, err := resp.Body.Read(buffer)
		if err == io.EOF || bytesRead == 0 {
			break
		}
		if err != nil {
			return -1, err
		}
		f.WriteAt(buffer[0:bytesRead], bytesWritten)
		bytesWritten += int64(bytesRead)
	}

	return int(bytesWritten), nil
}
