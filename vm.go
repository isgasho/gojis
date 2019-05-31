package vm

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/gojis/vm/runtime"
)

func Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("VM crashed: %v [recovered]\n", err)
			panic(err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gojis",
	Short: "Evaluates a set of .js files.",
	Args:  cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		return root(args...)
	},
}

func root(args ...string) error {
	r := runtime.New()

	err := r.LoadDirectory(args[0])
	if err != nil {
		return err
	}

	return r.Start()
}
