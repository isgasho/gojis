package vm

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
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
	RunE: func(c *cobra.Command, args []string) error {
		return root(args...)
	},
}

func root(args ...string) error {
	logLevel := zerolog.InfoLevel
	if verbose {
		logLevel = zerolog.DebugLevel
	}

	r := runtime.New(
		runtime.LogLevel(logLevel),
	)

	return r.Start()
}
