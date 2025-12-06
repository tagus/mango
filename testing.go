package mango

import (
	"testing"
	"time"
)

func VerifyUntilTimeout(t *testing.T, wait time.Duration, failedMsg string, fn func(t *testing.T) bool) {
	t.Helper()

	timeout := time.After(wait)
	for {
		select {
		case <-timeout:
			if fn(t) {
				return
			}
			t.Fatal(failedMsg)
			return
		default:
			if fn(t) {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}
