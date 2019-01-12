package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is overridden by link flags during build
var Version = "unset"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
