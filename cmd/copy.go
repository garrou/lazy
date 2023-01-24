package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	copyCmd = &cobra.Command{
		Use:   "copy [name] [newname]",
		Short: "Copy a lazy named [name] to [newname]",
		Long:  "Copy a lazy named [name] to a new lazy called [newname]",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, newName := lib.BindNames(args)
			err := lib.CopyLazy(oldName, newName)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(copyCmd)
}
