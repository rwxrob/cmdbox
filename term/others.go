// +build aix js nacl plan9 windows android solaris

package term

func init() {
	WinSize = winsize{80, 40, 100, 100} // complete fudge for lesser OSes
}
