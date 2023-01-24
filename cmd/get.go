package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get [name]",
		Short: "Get a lazy with name or all lazies",
		Long:  "Get a lazy named [name] or all lazies without name",
		Run: func(cmd *cobra.Command, args []string) {
			name := lib.BindOptionnalArg(args)
			data, err := lib.GetLazies(name)

			if err != nil {
				log.Fatal(err)
			}
			lib.Display(data)
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
