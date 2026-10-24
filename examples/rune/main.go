package main

import (
	"fmt"
	"os"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	val, err := chocobar.Rune(
		chocobar.Prompt("Enter a character: "),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(val)
}
