package main

import (
	"fmt"
	"os"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	val, err := chocobar.Byte(
		chocobar.Prompt("Enter byte: "),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(val)
}
