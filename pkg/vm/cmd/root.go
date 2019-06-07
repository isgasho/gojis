package cmd

import "github.com/spf13/cobra"

func Execute() error {
	return RootCmd.Execute()
}

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
