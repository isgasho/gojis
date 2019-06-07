package vm

import (
	"fmt"
	"os"
	"runtime/debug"

	"gitlab.com/gojis/vm/pkg/vm/cmd"
)

func Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("VM crashed: %v [recovered]\n\t%v\n", err, string(debug.Stack()))
		}
	}()

	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
