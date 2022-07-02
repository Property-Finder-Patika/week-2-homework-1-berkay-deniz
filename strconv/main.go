// Berkay Deniz

// An exercise for conversions among string and basic types using strconv package

package main

import (
	"fmt"
	"strconv"
)

func main() {

	birthDate := "1999"                         // Example string that represents a number
	convertedDate, e := strconv.Atoi(birthDate) // converts string into integer using Atoi function from strconv
	if e == nil {                               // Error check
		fmt.Printf("%T \n %v", convertedDate, convertedDate) // prints the number and its type
	}

}
