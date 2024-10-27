package mango

import (
	"log"
	"runtime/debug"
)

func Fatal(v ...any) {
	debug.PrintStack()
	log.Fatal(v...)
}

func FatalIf(err error, v ...any) {
	if err == nil {
		return
	}
	args := append([]any{err}, v...)
	Fatal(args...)
}
