package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove [name]",
		Short: "Remove a lazy",
		Long:  "Remove a lazy named [name]",
		Run: func(cmd *cobra.Command, args []string) {
			name := lib.BindArg(args)
			err := lib.RemoveLazy(name)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
