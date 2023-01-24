package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	copyCmd = &cobra.Command{
		Use:   "copy",
		Short: "copy [name] [newname]",
		Long:  "Copy a lazy named [name] to a new lazy called [newname]",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, newName, err := lib.BindNames(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if err := lib.CopyLazy(oldName, newName); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(copyCmd)
}
