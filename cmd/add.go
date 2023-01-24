package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add [NAME] [PATH]",
		Short: "Add a lazy with a name associated to a path",
		Long:  "Add a lazy with a name [NAME] associated to a path [PATH]",
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
