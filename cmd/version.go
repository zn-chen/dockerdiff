package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const version = "1.0.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Docker version information",
	Run: func(cmd *cobra.Command, args []string) {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", version)
	},
}
