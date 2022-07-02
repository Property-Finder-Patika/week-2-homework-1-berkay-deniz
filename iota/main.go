// Berkay Deniz

// An exercise for iota

package main

import (
	"fmt"
)

// A constant that keep some predefined colors as iota
const (
	Green int = iota
	Red
	Blue
	Yellow
	Orange
	Black
	White
)

func main() {

	// Print the values for iota colors
	fmt.Printf("The value of Green    is %v\n", Green)
	fmt.Printf("The value of Red is %v\n", Red)
	fmt.Printf("The value of Blue is %v\n", Blue)
	fmt.Printf("The value of Yellow  is %v\n", Yellow)
	fmt.Printf("The value of Orange   is %v\n", Orange)
	fmt.Printf("The value of Black is %v\n", Black)
	fmt.Printf("The value of White is %v\n", White)
}
