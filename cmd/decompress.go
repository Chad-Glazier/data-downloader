/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Chad-Glazier/fdd/decompress"
	"github.com/Chad-Glazier/fdd/misc"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// decompressCmd represents the decompress command
var decompressCmd = &cobra.Command{
	Use:   "decompress file",
	Short: "decompresses a file",
	Long: `
Infers the type of a local file and then decompresses it with an
appropriate algorithm. The decompressed copy will be written to the
current working directory. For example,

	fdd decompress ./file.zip

will extract the archive to ./file (a directory).

The exact algorithm used is based on the file's type, which is
inferred from the following properties:
  1) The file's extension. E.g., file name that ends with ".zip" is
     assumed to contain a zip archive.
  2) In lieu of a file extension, this command will look at the 
     file's "magic number"; i.e., the first few bytes of the file.
	 Most common compression methods will leave a signature so that
	 the type can be recognized.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		srcPath := args[0]

		src, err := os.Open(srcPath)
		if err != nil {
			log.Fatalf(
				"Error opening file %s:\n%s",
				srcPath,
				err.Error(),
			)
		}
		defer src.Close()

		// Generate a unique name for the extracted files. First, we try the
		// basename of the source. Next, we try adding the suffix "(1)", then
		// "(2)", and so on.
		destPath := misc.BaseName(srcPath)
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

		var fileSize int
		info, err := src.Stat()
		if err != nil {
			fileSize = -1
		} else {
			fileSize = int(info.Size())
		}

		fmt.Printf(
			"\nExtracting %s (%s)\n",
			srcPath,
			misc.ByteCountToString(fileSize),
		)

		progress := progressbar.NewOptions(
			fileSize,
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(20),
		)

		bytesWritten, err := decompress.File(srcPath, destPath, progress)
		if err != nil {
			log.Fatalf(
				"\nError extracting file:\n%s",
				err.Error(),
			)
		}

		cwd, err := os.Getwd()
		if err != nil {
			cwd = "./"
		}

		absoluteDest := filepath.Join(cwd, destPath)

		fmt.Printf(
			"\n\nFile extracted to %s (%s).\n\n", 
			absoluteDest, 
			misc.ByteCountToString(int(bytesWritten)),
		)
	},
}

func init() {
	rootCmd.AddCommand(decompressCmd)
}
