package runtime

import (
	"os"

	"github.com/rs/zerolog"
)

type Runtime struct {
	log zerolog.Logger
}

type RuntimeOption func(*Runtime)

func LogLevel(l zerolog.Level) RuntimeOption {
	return func(r *Runtime) {
		r.log = r.log.Level(l)
	}
}

func New(opts ...RuntimeOption) *Runtime {
	r := new(Runtime)

	r.log = zerolog.New(os.Stdout).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *Runtime) Start() error {
	panic("TODO")
}
