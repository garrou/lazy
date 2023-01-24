package cmd

import (
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:   "move",
		Short: "move [name] [newname]",
		Long:  "Move a lazy named [name] to a new lazy named [newname]",
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
