/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Chad-Glazier/fdd/decompress"
	"github.com/Chad-Glazier/fdd/filetypes"
	"github.com/Chad-Glazier/fdd/misc"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get url",
	Short: "download and decompress a file from a URL",
	Long: `Downloads the file at the given URL and decompresses it if necessary.
Uncompressed files will be downloaded like normal. For example,

	fdd get https://google.com/

will download the uncompressed Google homepage as an HTML file. A
compressed file would be treated differently; for example,

	fdd get https://some.website/with/a/file.zip

would determine that the file at the url is a compressed zip archive
and would download and decompress it to a directory named "file" in 
the working directory.

The exact type of a compressed file is guessed based on the following
properties:
  1) The file's extension. E.g., a URL whose path ends with ".zip"
     is assumed to yield a zip archive.
  2) The "Content-Type" header of the response. This property is only
     used when the file doesn't have an extension.`,
	Run: func(cmd *cobra.Command, args []string) {
		rawUrl := args[0]

		parsedUrl, err := url.Parse(rawUrl)
		if err != nil {
			log.Fatalf(
				"The URL provided %s is invalid:\n%s", rawUrl, err.Error(),
			)
		}

		resp, err := http.Get(rawUrl)
		if err != nil {
			log.Fatalf(
				"The request to %s failed:\n%s", rawUrl, err.Error(),
			)
		}

		if resp.StatusCode > 299 || resp.StatusCode < 200 {
			log.Fatalf(
				"The server sent back a bad response: %s", resp.Status,
			)
		}

		fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
		if err != nil {
			fileSize = -1
		}

		fileType := filetypes.MimeFromResponse(resp)

		fileName, _ := misc.FileNameFromUrl(rawUrl)

		if fileName == "" {
			fileName = strings.ReplaceAll(parsedUrl.Host, ".", "_")
		}

		if len(strings.Split(fileName, ".")) == 1 {
			// the filename doesnt have an extension
			fileName += filetypes.MimeToExt[fileType]
		}

		compressionAlgo := filetypes.MimeToCompression[fileType]

		if fileType != "" && fileSize != -1 {
			fmt.Printf(
				"\nRetrieving %s (%s, %s)...\n", 
				fileName,
				misc.ByteCountToString(fileSize),
				fileType,
			)
		} else if fileType != "" {
			fmt.Printf(
				"\nRetrieving %s (%s)...\n", 
				fileName,
				fileType,
			)
		} else if fileSize != -1 {
			fmt.Printf(
				"\nRetrieving %s (%s)...\n", 
				fileName,
				misc.ByteCountToString(fileSize),
			)
		} else {
			fmt.Printf(
				"\nRetrieving %s...\n", 
				fileName,
			)
		}

		progress := progressbar.NewOptions(
			fileSize,
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(20),
		)

		cwd, err := os.Getwd()
		if err != nil {
			cwd = "./"
		}

		r := resp.Body
		defer r.Close()

		// if the file is uncompressed
		if compressionAlgo == "" {
			w, err := os.Create(fileName)
			if err != nil {
				log.Fatalf("Error creating the file %s:\n%s", fileName, err.Error())
			}
			defer w.Close()

			bytesWritten, err := io.Copy(io.MultiWriter(w, progress), r)
			if err != nil {
				log.Fatalf("\nError downloading data:\n%s", err.Error())
			}

			absoluteDest := filepath.Join(cwd, fileName)
	
			if fileSize == -1 {
				fmt.Printf(
					"\n\nFile saved to %s (%s)\n\n",
					absoluteDest, 
					misc.ByteCountToString(int(bytesWritten)),
				)
			} else {
				fmt.Printf(
					"\n\nFile saved to %s\n\n",
					absoluteDest,
				)
			}

			return
		}

		// Generate a unique name for the extracted files. First, we try the
		// basename of the source. Next, we try adding the suffix "(1)", then
		// "(2)", and so on.
		destPath := misc.BaseName(fileName)
		for i := range math.MaxInt {
			suffix := " (" + strconv.Itoa(i) + ")"
			if i == 0 {
				suffix = ""
			}

			_, err := os.Stat(destPath + suffix)
			if os.IsNotExist(err) {
				destPath += suffix
				break
			}
		}
		absoluteDest := filepath.Join(cwd, destPath)

		bytesWritten, err := decompress.General(compressionAlgo, r, destPath, progress)
		if err != nil {
			log.Fatalf("\nError extracting the file:\n%s", err.Error())
		}

		fmt.Printf(
			"\n\nFile downloaded and decompressed to %s (%s)\n\n",
			absoluteDest,
			misc.ByteCountToString(int(bytesWritten)),
		)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
