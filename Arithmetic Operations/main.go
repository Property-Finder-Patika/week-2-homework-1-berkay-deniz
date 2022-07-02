// Berkay Deniz

// An exercise for arithmetic operations
package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	var operator string

	fmt.Print("Enter First number: ")
	fmt.Scanln(&firstNumber) // Gets the first input
	fmt.Print("Enter Second number: ")
	fmt.Scanln(&secondNumber)                // Gets the second input
	fmt.Print("Enter Operator (+,-,/,%,*):") // Gets the operator as input
	fmt.Scanln(&operator)
	output := 0
	switch operator { // switch-case structure for one case for each operation
	case "+":
		output = firstNumber + secondNumber
		break
	case "-":
		output = firstNumber - secondNumber
	case "/":
		output = firstNumber / secondNumber
	case "*":
		output = firstNumber * secondNumber
	case "%":
		output = firstNumber % secondNumber
	default:
		fmt.Println("Invalid Operator") // Default case for unknown operators
	}

	// makes the calculation and prints as formatted
	fmt.Printf("%d %s %d = %d", firstNumber, operator, secondNumber, output)
}
