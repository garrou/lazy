package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	copyCmd = &cobra.Command{
		Use:   "cp [NAME] [NEWNAME]",
		Short: "Copy a lazy",
		Long:  "Copy a lazy named [NAME] to a new lazy named [NEWNAME]",
		Run: func(cmd *cobra.Command, args []string) {
			err := lib.CheckArgs(args, 2)

			if err != nil {
				fmt.Println(err)
				cmd.Usage()
			} else if err := lib.CopyLazy(args[0], args[1]); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(copyCmd)
}
