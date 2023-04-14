package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	removeCmd = &cobra.Command{
		Use:   "rm [NAME]",
		Short: "Remove a lazy",
		Long:  "Remove a lazy named [NAME]",
		Run: func(cmd *cobra.Command, args []string) {
			err := lib.CheckArgs(args, 1)

			if err != nil {
				fmt.Println(err)
				cmd.Usage()
			} else if err := lib.RemoveLazy(args[0]); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(removeCmd)
}
