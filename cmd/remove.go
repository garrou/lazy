package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove [NAME]",
		Short: "Remove a lazy",
		Long:  "Remove a lazy named [NAME]",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := lib.BindName(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if err := lib.RemoveLazy(name); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
