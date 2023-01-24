package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	printCmd = &cobra.Command{
		Use:   "print",
		Short: "print",
		Long:  "Print all lazies",
		Run: func(cmd *cobra.Command, args []string) {
			if data, err := lib.GetLazies(); err != nil {
				log.Fatal(err)
			} else {
				lib.Display(data)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(printCmd)
}
