package isatty

import "golang.org/x/sys/unix"

func isattyIntl(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	return err == nil
}
