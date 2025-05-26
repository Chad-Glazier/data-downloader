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

const VERSION = "development version 0.0.2"

func main() {
	if os.Args[1] == "version" {
		fmt.Println(VERSION)
		return
	}

	args, err := parseArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := http.Get(args.url)
	if err != nil {
		fmt.Printf("The URL %s sent an invalid response.\n", args.url)
		return
	}
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		fmt.Printf(
			"The URL %s sent a bad response: %s\n", 
			args.url, resp.Status,
		)
	}

	
	fileType := filetypes.MimeFromName(args.output)
	inferred := true
	if (fileType == "") {
		fileType = filetypes.MimeFromResponse(resp)
		inferred = false
	}

	if args.output == "" {
		args.output = "data" + filetypes.MimeToExt[fileType]
	}

	fmt.Printf("Found file %s\n", args.output)

	fmt.Printf("  - file type: %s", fileType)
	if inferred {
		fmt.Printf(" (inferred from file extension)\n")
	} else {
		fmt.Printf(" (according to server response)\n")
	}

	fileSize := resp.ContentLength
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
			fmt.Printf("Downloading file to %s\n", args.output)
		} else {
			fmt.Printf("Downloading file to %s\n", filepath.Join(cwd, args.output))
		}

		fileSize, err := writeBodyToFile(args.output, resp)
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
		fmt.Printf("Downloading file to %s\n", misc.BaseName(args.output))
	} else {
		fmt.Printf(
			"Downloading file to %s\n",
			filepath.Join(cwd, misc.BaseName(args.output)),
		)
	}

	bytesWritten, err := decompress.General(
		compressionAlgo, 
		resp.Body, 
		misc.BaseName(args.output),
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

