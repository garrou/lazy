package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	replaceCmd = &cobra.Command{
		Use:   "replace",
		Short: "replace [name] [path]",
		Long:  "Replace a lazy named [name] with the new path [path]",
		Run: func(cmd *cobra.Command, args []string) {
			name, path, err := lib.BindNamePath(args)

			if err != nil {
				fmt.Println(cmd.Usage())
				return
			}
			if err := lib.ReplaceLazy(name, path); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(replaceCmd)
}
