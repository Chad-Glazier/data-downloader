package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Chad-Glazier/data-downloader/decompress"
	"github.com/Chad-Glazier/data-downloader/filetypes"
	"github.com/Chad-Glazier/data-downloader/misc"
)

const VERSION = "development version 0.0.3"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("this command expects at least one argument.")
	}

	if os.Args[1] == "version" {
		fmt.Println(VERSION)
		return
	}

	downloadUrl := os.Args[1]

	resp, err := http.Get(downloadUrl)
	if err != nil {
		fmt.Printf("The URL %s sent an invalid response.\n", downloadUrl)
		return
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		fmt.Printf(
			"The URL %s sent a bad response: %s\n", 
			downloadUrl, resp.Status,
		)
	}

	filename, _ := misc.FileNameFromUrl(resp.Request.URL.String())
	fileType := filetypes.MimeFromResponse(resp)
	if filename == "" {
		filename = "data"
	}
	if fileType != "" && !misc.HasExtension(filename) {
		filename += filetypes.MimeToExt[fileType]
	}
	fileSize := resp.ContentLength

	fmt.Printf("Found file %s\n", filename)
	fmt.Printf("  - file type: %s\n", fileType)
	if fileSize == -1 {
		fmt.Printf("  - file size unknown\n")
	} else {
		fmt.Printf("  - file size: %s\n", misc.ByteCountToString(int(fileSize)))
	}

	compressionAlgo, ok := filetypes.MimeToCompression[fileType]
	if !ok {
		// file is uncompressed
		fmt.Printf("  - file is uncompressed\n")
		
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Downloading file to %s\n", filename)
		} else {
			fmt.Printf("Downloading file to %s\n", filepath.Join(cwd, filename))
		}

		fileSize, err := misc.WriteBodyToFile(filename, resp)
		if err != nil {
			fmt.Println("Error downloading file.")
		} else {
			fmt.Printf("Download complete (%s)", misc.ByteCountToString(fileSize))
		}

		return
	}

	compressionAlgoName := strings.Split(compressionAlgo, "/")[1]
	fmt.Printf("  - file is compressed with %s\n", compressionAlgoName)

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Downloading file to %s\n", misc.BaseName(filename))
	} else {
		fmt.Printf(
			"Downloading file to %s\n",
			filepath.Join(cwd, misc.BaseName(filename)),
		)
	}

	bytesWritten, err := decompress.General(
		compressionAlgo, 
		resp.Body, 
		misc.BaseName(filename),
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf(
		"Download complete (%s, uncompressed)",
		misc.ByteCountToString(bytesWritten),
	)
}
