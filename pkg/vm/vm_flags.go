package vm

import "gitlab.com/gojis/vm/internal/runtime"

var (
	verbose bool
)

func init() {
	flags := rootCmd.PersistentFlags()

	flags.BoolVarP(&runtime.Debug, "debug", "d", false, "Enables debug mode (will impact performance, as it also will not skip log messages)")
	flags.IntVar(&runtime.LogBufferSize, "log-buffer", 2000, "Sets the buffer size of the logger. This value should be increased if the logger drops many messages or the log suggests to because of diode collisions.")
}
