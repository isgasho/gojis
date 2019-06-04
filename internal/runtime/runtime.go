package runtime

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-multierror"
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

// LoadDirectory recursively loads all files in a given directory
// with LoadFile. After loading all files, an error containing
// all occurred errors will be returned.
// If no errors occurred, nil will be returned.
func (r *Runtime) LoadDirectory(path string) error {
	var result *multierror.Error

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		result = multierror.Append(result, err)

		if info.IsDir() {
			return nil // skip if it's a directory
		}

		err = r.LoadFile(path)
		if err != nil {
			result = multierror.Append(result, err)
		}

		return nil // do not stop processing the directory recursively
	})

	return result.ErrorOrNil()
}

// LoadFiles loads all given paths with LoadFile.
// After loading all files, an error containing
// all occurred errors will be returned.
// If no errors occurred, nil will be returned.
func (r *Runtime) LoadFiles(paths ...string) error {
	var result *multierror.Error

	for _, path := range paths {
		err := r.LoadFile(path)
		if err != nil {
			result = multierror.Append(result, err)
		}
	}

	return result.ErrorOrNil()
}

// LoadFile parses all code in a given JavaScript file.
// If the file is not a JavaScript file, it will be skipped silently.
// The parsed AST will be used upon code execution.
func (r *Runtime) LoadFile(path string) error {
	if !IsJavaScriptFile(path) {
		return nil
	}

	r.log.Debug().
		Str("file", path).
		Msg("load file")

	err := r.parser.ParseFile(path)
	if err != nil {
		return err
	}

	return nil
}

// IsJavaScriptFile returns true if and only if the
// extension of a given file path is '.js'.
func IsJavaScriptFile(path string) bool {
	return filepath.Ext(path) == ".js"
}

func (r *Runtime) Start() error {
	return nil
}
