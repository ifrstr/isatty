package isatty

import "golang.org/x/sys/windows"

func isattyIntl(fd uintptr) bool {
	var mode uint32
	err := windows.GetConsoleMode(windows.Handle(fd), &mode)
	return err == nil
}
