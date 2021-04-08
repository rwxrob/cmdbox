package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ArgsFromStdin converts each line of standard input into a slice of
// strings suitable for using as arguments. If fields is passed any
// occurance of {n} will be replaced with the appropriate field in order
// with n beginning at 1.
func ArgsFromStdin(fields ...string) []string {
	args := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for n, v := range fields {
			lk := fmt.Sprintf("{%v}", n+1)
			line = strings.ReplaceAll(line, lk, v)
		}
		args = append(args, line)
	}
	return args
}
