package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	replaceCmd = &cobra.Command{
		Use:   "replace [name] [path]",
		Short: "Replace a lazy named [name] with [path]",
		Long:  "Replace a lazy named [name] with the new path [path]",
		Run: func(cmd *cobra.Command, args []string) {
			name, path := lib.BindNamePath(args)
			err := lib.ReplaceLazy(name, path)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(replaceCmd)
}
