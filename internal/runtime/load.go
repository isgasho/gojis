package runtime

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

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
