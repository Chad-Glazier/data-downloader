package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Chad-Glazier/fdd/filetypes"
	"github.com/Chad-Glazier/fdd/misc"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download url",
	Short: "download a file from a URL without decompressing it",
	Long: `
Downloads a file from a given URL into the current working directory.
For example,

	fdd download https://some.domain/file.zip/

will download the response to ./file.zip, leaving it compressed.
`,
	Args: cobra.ExactArgs(1),
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

		w, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("Error creating the file %s:\n%s", fileName, err.Error())
		}
		defer w.Close()

		r := resp.Body
		defer r.Close()

		if fileType != "" && fileSize != -1 {
			fmt.Printf(
				"\nDownloading %s (%s, %s)...\n", 
				fileName,
				misc.ByteCountToString(fileSize),
				fileType,
			)
		} else if fileType != "" {
			fmt.Printf(
				"\nDownloading %s (%s)...\n", 
				fileName,
				fileType,
			)
		} else if fileSize != -1 {
			fmt.Printf(
				"\nDownloading %s (%s)...\n", 
				fileName,
				misc.ByteCountToString(fileSize),
			)
		} else {
			fmt.Printf(
				"\nDownloading %s...\n", 
				fileName,
			)
		}

		progress := progressbar.NewOptions(
			fileSize,
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(20),
		)

		bytesWritten, err := io.Copy(io.MultiWriter(w, progress), r)
		if err != nil {
			log.Fatalf("\nError downloading data:\n%s", err.Error())
		}

		cwd, err := os.Getwd()
		if err != nil {
			cwd = "./"
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
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
