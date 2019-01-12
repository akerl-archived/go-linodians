package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/akerl/go-linodians/api"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff OLD NEW",
	Short: "diff employees",
	RunE:  diffRunner,
}

func init() {
	rootCmd.AddCommand(diffCmd)
}

func diffRunner(_ *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("too few args provided. Check --help for more info")
	}
	oldC, err := loadJSON(args[0])
	if err != nil {
		return err
	}
	newC, err := loadJSON(args[1])
	if err != nil {
		return err
	}

	ds := api.Diff(oldC, newC)

	staticFunc := func(e api.Employee) { fmt.Printf("  %s -- %s -- %s\n", e.Username, e.Fullname, e.Title) }
	changeFunc := func(e api.Employee) {
		oe := oldC[e.Username]
		fmt.Printf("  %s -- %s -- %s -> %s\n", e.Username, e.Fullname, oe.Title, e.Title)
	}

	printDiff("Added", ds.Added, staticFunc)
	printDiff("Modified", ds.Modified, changeFunc)
	printDiff("Removed", ds.Removed, staticFunc)

	return nil
}

func printDiff(title string, c api.Company, f func(api.Employee)) {
	if len(c) > 0 {
		fmt.Printf("%s:\n", title)
		for _, e := range c {
			f(e)
		}
	}
}

func loadJSON(file string) (api.Company, error) {
	var c api.Company

	data, err := ioutil.ReadFile(file) // #nosec
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(data, &c)
	return c, err
}
