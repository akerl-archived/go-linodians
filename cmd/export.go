package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/akerl/go-linodians/api"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [FILE]",
	Short: "export employees",
	RunE:  exportRunner,
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

func exportRunner(cmd *cobra.Command, args []string) error {
	list, err := api.Load()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Println(string(data))
	} else {
		filename := args[0]
		err = ioutil.WriteFile(filename, data, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
