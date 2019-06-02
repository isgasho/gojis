package vm

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
	"gitlab.com/gojis/vm/internal/runtime"
)

func Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("VM crashed: %v [recovered]\n\t%v\n", err, string(debug.Stack()))
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
	Run: func(c *cobra.Command, args []string) {
		root(args...)
	},
}

func root(args ...string) {
	r := runtime.New()

	err := r.LoadDirectory(args[0])
	if err != nil {
		fmt.Printf("Error occurred while loading files: %v\n", err)
	}

	err = r.Start()
	if err != nil {
		panic(err)
	}
}
