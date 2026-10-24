package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	price, err := chocobar.Float64(
		chocobar.Prompt("Enter price: "),
		chocobar.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed to read price: %v", err)
	}
	fmt.Printf("Price entered: %.2f\n", price)
}
