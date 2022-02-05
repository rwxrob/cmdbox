package util_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleGetSection() {
	r := strings.NewReader(`

# start
some
thing
# stop

	`)

	fmt.Println(util.GetSection(r, "# start", "# stop"))

	// Output:
	// some
	// thing
}
