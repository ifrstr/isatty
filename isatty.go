// Package isatty determine whether the stream is a tty.
//
// # Features
//
// - Supports most platforms
// - Supports force override using environment variable `FORCE_TTY`
//
// # Usage
//
//	import "gopkg.ilharper.com/x/isatty"
//
//	if isatty.Isatty(os.Stdout.Fd()) {
//	    // TTY
//	} else {
//	    // Not TTY
//	}
//
// # Using `FORCE_TTY`
//
// User can override the result using the environment variable `FORCE_TTY`.
//
//	# Force TTY
//	FORCE_TTY=
//	FORCE_TTY=1
//	FORCE_TTY=true
//
//	# Force not TTY
//	FORCE_TTY=0
//	FORCE_TTY=false
package isatty

import (
	"os"
	"strconv"
	"strings"
)

func Isatty(fd uintptr) bool {
	env := os.Environ()

	for _, e := range env {
		if strings.HasPrefix(e, "FORCE_TTY=") {
			s := e[10:]

			if len(s) == 0 {
				return true
			}

			if s == "true" {
				return true
			}

			if s == "false" {
				return false
			}

			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return true
			}

			return i > 0
		}
	}

	return isattyIntl(fd)
}
