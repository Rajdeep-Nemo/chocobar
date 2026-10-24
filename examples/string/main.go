package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	name, err := chocobar.String(
		chocobar.Prompt("Enter username: "),
	)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	fmt.Printf("Username: %s\n", name)
}
