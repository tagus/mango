package mango

import (
	"fmt"
	"os"
	"runtime/debug"
)

func FatalIf(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Fatal(err error) {
	fmt.Println(err)
	debug.PrintStack()
	os.Exit(1)
}
