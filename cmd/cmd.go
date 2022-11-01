package main

import (
	"github.com/spf13/cobra"
)

var output string

func init() {
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Write to a file, instead of STDOUT")
}

var rootCmd = &cobra.Command{
	Use:     "dockerdiff save IMAGE1 IMAGE2",
	Example: "dockerdiff save IMAGE1 IMAGE2 -o dst.tar",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
