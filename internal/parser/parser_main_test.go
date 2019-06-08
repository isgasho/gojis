package parser

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

const (
	parserTestRepo = "https://github.com/tc39/test262-parser-tests"
)

func TestMain(m *testing.M) {
	setup()
	defer tearDown()

	os.Exit(m.Run())
}

func setup() {
	cloneParserTestRepo()
}

func tearDown() {}

func cloneParserTestRepo() {
	if testing.Short() {
		// don't clone if short testing
		return
	}

	if _, err := os.Stat("test262-parser-tests"); os.IsNotExist(err) {
		log.Println("Parser test directory does not exist, cloning it...")

		var stdout, stderr bytes.Buffer

		cmd := exec.Command("git", "clone", parserTestRepo)
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			panic(fmt.Sprintf("Clone failed.\nStdout: '%v'\nStderr: '%v'", stdout.String(), stderr.String()))
		}
	}
}
