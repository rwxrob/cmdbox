package util_test

import (
	"fmt"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleEmphasize() {
	fmt.Println(util.Emphasize("**foo**"))
}
