package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get [name]",
		Long:  "Get a lazy named [name]",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := lib.BindName(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if data, err := lib.GetLazy(name); err != nil {
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
