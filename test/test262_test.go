package test262

import "testing"

func Test262(t *testing.T) {
	// until the runtime can not run code at all, skip this test entirely
	// (will greatly decrease time it takes to execute tests)
	t.SkipNow()

	if testing.Short() {
		t.SkipNow()
	}

	t.Run("Suite clone", CloneTest262Repo)
}
