package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	temp, err := chocobar.Float32(
		chocobar.Prompt("Enter temperature: "),
		chocobar.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read temperature: %v", err)
	}
	fmt.Printf("Temperature entered: %.2f\n", temp)
}
