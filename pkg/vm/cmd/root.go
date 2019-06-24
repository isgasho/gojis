package cmd

import "github.com/spf13/cobra"

// Execute is used to execute the RootCmd.
func Execute() error {
	return RootCmd.Execute()
}

// RootCmd is the root cobra command of the VM.
var RootCmd = &cobra.Command{
	Use:   "gojis",
	Short: "Evaluates a set of .js files.",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		root(args...)
	},
}

func root(args ...string) {
	panic("TODO")
}
