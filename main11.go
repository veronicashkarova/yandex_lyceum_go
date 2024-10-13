package main
import (
	"fmt"
	"log"
)
// Recursive approach to calculate factorial
func factorialRecursive(n int) int {
	// Base case: factorial of 0 or 1 is 1
	if n == 0 || n == 1 {
		return 1
	}
	// Recursive call to calculate factorial
	return n * factorialRecursive(n-1)
}
func main() {
	// Calculate factorial using recursive approach
	var number int
	_, err := fmt.Scanln(&number)
	if err != nil {
		log.Fatal("Invalid input. Please enter a valid non-negative integer.")
	}
	if number < 0 {
		log.Fatal("Invalid input. Please enter a non-negative integer.")
	}
	// Calculate factorial using recursive approach
	factorialRec := factorialRecursive(number)
	fmt.Println(factorialRec)
}