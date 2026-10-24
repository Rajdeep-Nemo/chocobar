package main

import (
	"fmt"
	"os"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	val, err := chocobar.Int8(
		chocobar.Prompt("Enter a number: "),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(val)
}
