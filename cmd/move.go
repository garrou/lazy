package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:   "move [NAME] [NEWNAME]",
		Short: "Move a lazy",
		Long:  "Move a lazy named [NAME] to a new lazy named [NEWNAME]",
		Run: func(cmd *cobra.Command, args []string) {
			oldName, newName, err := lib.BindNames(args)

			if err != nil {
				cmd.Usage()
				return
			}
			if err := lib.MoveLazy(oldName, newName); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(moveCmd)
}
