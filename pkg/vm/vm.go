package vm

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/TimSatke/gojis/pkg/vm/cmd"
)

// Run is the entry point of the VM.
// It will recover any panic and print its message, including a stack trace.
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
