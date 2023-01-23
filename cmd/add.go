package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add [name] [path]",
		Short: "Add a lazy",
		Long:  "Add a lazy named [name] associated to path [path]",
		Run: func(cmd *cobra.Command, args []string) {
			name, path := lib.Bind2Args(args)
			err := lib.AddLazy(name, path)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
