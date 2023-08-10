package mango

import (
	"os"
	"runtime/debug"
)

func FatalIf(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Fatal(err error) {
	debug.PrintStack()
	os.Exit(1)
}
