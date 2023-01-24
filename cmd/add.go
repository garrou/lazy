package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add [name] [path]",
		Long:  "Add a lazy named [name] associated to path [path]",
		Run: func(cmd *cobra.Command, args []string) {
			name, path, err := lib.BindNamePath(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if err := lib.AddLazy(name, path); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
