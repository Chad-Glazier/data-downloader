/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fdd",
	Short: "A utility for downloading and decompressing files",
	Long: `A command-line tool for downloading and decompressing files from URLs
or from the local machine, inferring types in order to  automatically
determine an appropriate decompression algorithm. For example:

	fdd get https://google.com/

This will determine that the server's response is an HTML file, and
will download it with a default name.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
}

