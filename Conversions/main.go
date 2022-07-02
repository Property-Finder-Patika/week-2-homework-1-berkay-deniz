// Berkay Deniz

// An exercise for conversions among basic types

package main

import (
	"fmt"
	"math"
)

func main() {

	var edge1, edge2 int = 8, 15 // Rpresent the two edges of a right triangle

	// Find the hypotenuse edge by using Pythagoras' theorem
	var hypotenuse float64 = math.Sqrt(float64(edge1*edge1 + edge2*edge2)) // convert to float
	var edge3 uint = uint(hypotenuse)                                      // convert to uint
	fmt.Println(edge1, edge2, edge3)                                       // print results
}
