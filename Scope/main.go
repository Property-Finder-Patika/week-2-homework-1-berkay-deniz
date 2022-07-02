// Berkay Deniz

// An exercise for scope

package main

import "fmt"

// global variable
var g int

func main() {

	// local variables
	var a, b int

	// initializing all local and global variables
	a = 10
	b = 20
	g = a + b

	fmt.Printf("Values： a = %d, b = %d and g = %d\n", a, b, g)
}
