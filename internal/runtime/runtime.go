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
	LoadBuffer    = 100
)

type Runtime struct {
	log zerolog.Logger

	parser *parser.Parser
}

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

func debugLogger(w io.Writer) zerolog.Logger {
	return zerolog.New(w).With().
		Timestamp().
		// Caller().
		Logger().
		Level(zerolog.DebugLevel)
}

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
