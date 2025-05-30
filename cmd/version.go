package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "displays the current version of fdd",
	Long: `Displays the current version of fdd.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
