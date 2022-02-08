package util_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/cmdbox/util"
)

func ExampleFetch() {
	//defer os.Remove("testdata/fetched")
	util.Fetch("https://rwx.gg/README.md", "testdata/fetched", 10)
	fetched, err := os.ReadFile("testdata/fetched")
	if err != nil {
		fmt.Println(err)
	}
	if len(fetched) > 0 {
		fmt.Println("something")
	}
	// Output:
	// something
}
