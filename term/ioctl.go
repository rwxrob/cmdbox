// +build !aix
// +build !js
// +build !nacl
// +build !plan9
// +build !windows
// +build !android
// +build !solaris

package term

import (
	"syscall"
	"unsafe"
)

func init() {
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(0), uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&WinSize)))
}
