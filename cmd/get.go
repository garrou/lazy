package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get [NAME]",
		Short: "Get a lazy",
		Long:  "Get a lazy named [NAME]",
		Run: func(cmd *cobra.Command, args []string) {
			err := lib.CheckArgs(args, 1)

			if err != nil {
				fmt.Println(err)
				cmd.Usage()
				return
			}
			if data, err := lib.GetLazy(args[0]); err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(data)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
