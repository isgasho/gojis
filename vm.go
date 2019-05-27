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
			fmt.Printf("VM crashed: %v\n", err)
			os.Exit(1)
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
	Run: func(_ *cobra.Command, args []string) {
		root(args...)
	},
}

func root(args ...string) {
	r := runtime.New()
	r.SaySomething()
}
