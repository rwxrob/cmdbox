package term

import (
	"os"
)

// WinSize is 80x24 by default but is detected and set to a more
// accurate value at init() time on systems that support ioctl
// (currently). This value can be overriden by those wishing a more
// consistent value or who prefer not to fill the screen completely when
// displaying help and usage information.
var WinSize = &struct {
	Row, Col       uint16
	Xpixel, Ypixel uint16
}{80, 24, 100, 100}

// ForceYes forces IsTerminal to always return true. If ForceNo is also
// set the cumulative effect is negated (as if neither were set). See
// ForceNo.
var ForceYes bool

// ForceNo forces IsTerminal to always return false. See ForceYes.
var ForceNo bool

// IsTerminal returns true if the output is to an interactive terminal
// (not piped in any way). This is useful when detemining if an extra
// line return is needed to avoid making programs chomp the line returns
// unnecessarily.
func IsTerminal() bool {
	switch {
	case ForceYes && !ForceNo:
		return true
	case ForceNo && !ForceYes:
		return false
	}
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}
