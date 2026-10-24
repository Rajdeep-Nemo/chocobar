package main

import (
	"fmt"
	"log"

	"github.com/Rajdeep-Nemo/chocobar"
)

func main() {
	// Example 1: Basic integer input
	age, err := chocobar.Int(
		chocobar.Prompt("Enter your age: "),
	)
	if err != nil {
		log.Fatalf("Failed to read age: %v", err)
	}
	fmt.Printf("Your age is %d.\n\n", age)

	// Example 2: Retrying up to 3 times on invalid input
	pin, err := chocobar.Int(
		chocobar.Prompt("Enter a number (max 3 attempts): "),
		chocobar.MaxAttempt(3),
	)
	if err != nil {
		log.Fatalf("Failed after 3 attempts: %v", err)
	}
	fmt.Printf("Number accepted: %d\n\n", pin)

	// Example 3: Infinite retries until a valid integer is provided
	score, err := chocobar.Int(
		chocobar.Prompt("Enter game score (retries until valid): "),
		chocobar.MaxAttempt(0),
	)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Final Score recorded: %d\n\n", score)
}
