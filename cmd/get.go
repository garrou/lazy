package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get [name]",
		Long:  "Get a lazy named [name] or all lazies without name",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := lib.BindOptionalName(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if data, err := lib.GetLazies(name); err != nil {
				log.Fatal(err)
			} else {
				lib.Display(data)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
