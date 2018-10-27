package cmd

import (
	"fmt"

	"github.com/akerl/go-linodians/api"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List employees",
	RunE:  listRunner,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listRunner(cmd *cobra.Command, args []string) error {
	list, err := api.Load()
	if err != nil {
		return err
	}
	for _, x := range list {
		fmt.Printf("%s -- %s -- %s\n", x.Username, x.Fullname, x.Title)
	}
	return nil
}
