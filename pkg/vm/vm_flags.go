package vm

import "github.com/TimSatke/gojis/pkg/vm/cmd"

var (
	Debug         bool
	LogBufferSize int
)

func init() {
	flags := cmd.RootCmd.PersistentFlags()

	flags.BoolVarP(&Debug, "debug", "d", false, "Enables debug mode (will impact performance, as it also will not skip log messages)")
	flags.IntVar(&LogBufferSize, "log-buffer", 2000, "Sets the buffer size of the logger. This value should be increased if the logger drops many messages or the log suggests to because of diode collisions.")
}
