package main

import (
	"fmt"
	"os"

	"github.com/Rajdeep-Nemo/sleepycat"
)

func main() {
	val, err := sleepycat.Uint16(
		sleepycat.Prompt("Enter a number: "),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(val)
}
