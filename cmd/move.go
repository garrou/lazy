package cmd

import (
	"fmt"
	"lazy/lib"
	"log"

	"github.com/spf13/cobra"
)

var (
	moveCmd = &cobra.Command{
		Use:   "mv [NAME] [NEWNAME]",
		Short: "Move a lazy",
		Long:  "Move a lazy named [NAME] to a new lazy named [NEWNAME]",
		Run: func(cmd *cobra.Command, args []string) {
			err := lib.CheckArgs(args, 2)

			if err != nil {
				fmt.Println(err)
				cmd.Usage()
			} else if err := lib.MoveLazy(args[0], args[1]); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(moveCmd)
}
