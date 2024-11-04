package main

import "fmt"

// Function to add two integers
func Ajouter(a int, b int) int {
	return a + b
}

// Function to multiply two integers
func multiplication(a int, b int) int {
	return a * b
}

// Function to divide two integers
func division(a int, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}
