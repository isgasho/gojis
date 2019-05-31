package runtime

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

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

func IsJavaScriptFile(path string) bool {
	return filepath.Ext(path) == ".js"
}
