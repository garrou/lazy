package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add [name] [path]",
		Short: "Add a path",
		Long:  `Add a shortcut [path] with the [name]`,
		Run: func(cmd *cobra.Command, args []string) {
			name, path := lib.BindArgs(args)
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
