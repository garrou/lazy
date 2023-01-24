package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:   "move [name] [newname]",
		Short: "Move a lazy named [name] to [newname]",
		Long:  "Move a lazy named [name] to a new lazy called [newname]",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, newName := lib.BindNames(args)
			err := lib.MoveLazy(oldName, newName)

			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(moveCmd)
}
