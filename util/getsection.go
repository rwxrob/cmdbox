package util

import (
	"bufio"
	"io"
	"strings"
)

// GetSection will return the content of all lines between the beginning
// and ending lines indicated. The match must be exact and all potential
// carriage returns in the section will be discarded.
//
func GetSection(src io.Reader, beg, end string) string {
	buf := ""
	insection := false
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, beg) {
			insection = true
			continue
		}
		if insection && strings.HasPrefix(line, end) {
			return buf
		}
		if insection {
			buf += scanner.Text() + "\n"
		}
	}
	return buf
}
