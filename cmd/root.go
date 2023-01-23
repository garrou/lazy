package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lazy",
	Short: "lazy helps you to retrieve your paths",
	Long: `A tool to manage your paths and access them easily.
			Complete documentation is available at https://github.com/garrou/lazy`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
