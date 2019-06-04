package runtime

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"gitlab.com/gojis/vm/internal/parser"
)

// Flags that are set, internal default values are set, external
// ones may differ
var (
	Debug         = false
	LogBufferSize = 2000
)

// Runtime is the core of the vm.
// The runtime is used to load and run
// source files.
//
// After creating a runtime object,
// use LoadFile(s) and/or LoadDirectory
// to load source files.
// Once source files have been loaded,
// execution can be started with Start.
type Runtime struct {
	log zerolog.Logger

	parser *parser.Parser
}

// New creates a new runtime object
// with a default configuration.
// Any runtime created with this
// function is ready to use.
func New() *Runtime {
	r := new(Runtime)

	w := os.Stdout

	if Debug {
		r.log = debugLogger(w)
	} else {
		r.log = defaultLogger(w)
	}

	r.parser = parser.New()

	return r
}

// debugLogger creates a blocking logger that will
// print every message.
// The logger's level is DEBUG.
func debugLogger(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().
		Timestamp().
		// Caller().
		Logger().
		Level(zerolog.DebugLevel)
}

// defaultLogger creates a non-blocking logger
// with a message buffer size of LogBufferSize.
func defaultLogger(w io.Writer) zerolog.Logger {
	wr := diode.NewWriter(w, LogBufferSize, 5*time.Millisecond, func(missed int) {
		fmt.Printf("Logger dropped %d messages\n", missed)
	})

	log := zerolog.New(wr).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)
	return log
}
