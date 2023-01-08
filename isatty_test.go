package isatty_test

import (
	"os"
	"testing"

	"gopkg.ilharper.com/x/isatty"
)

func Test(t *testing.T) {
	if isatty.Isatty(os.Stdout.Fd()) {
		println("os.Stdout is a TTY.")
	} else {
		println("os.Stdout is not a TTY.")
	}
}
