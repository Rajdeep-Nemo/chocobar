package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	enabled, err := chocobar.Bool(
		chocobar.Prompt("Enable feature? "),
		chocobar.MaxAttempt(3),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enabled:", enabled)
}
