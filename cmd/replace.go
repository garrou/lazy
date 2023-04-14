package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	replaceCmd = &cobra.Command{
		Use:   "rp [NAME] [PATH]",
		Short: "Replace the path of lazy",
		Long:  "Replace the path of lazy named [NAME] with the path [PATH]",
		Run: func(cmd *cobra.Command, args []string) {
			name, path, err := lib.CheckArgsNamePath(args)

			if err != nil {
				fmt.Println(err)
				cmd.Usage()
			} else if err := lib.ReplaceLazy(name, path); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(replaceCmd)
}
