package util

import (
	"bufio"
	"io"
	"strings"
)

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
