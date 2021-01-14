package main

import "fmt"

func main() {
	// here is an example of a variable
	name := "Maria"
	fmt.Println(name)

	// Data types
	// Example of a String
	var ourString string
	ourString = "wassup"
	fmt.Printf("Type: %T\n", ourString)

	// int
	var ourInt int
	ourInt = 55
	fmt.Printf("Type: %T\n", ourInt)

	// bool
	var ourBool bool
	ourBool = false
	fmt.Printf("Type: %T\n", ourBool)

	// float
	var ourFloat float32
	ourFloat = 5.55878
	fmt.Printf("Type: %T\n", ourFloat)

	// slice
	var ourSlice []string
	ourSlice = []string{"hello", "goodbye", "Aaron"}
	fmt.Println(ourSlice)

	// rune byte
	var ourRune rune
	ourRune = 'h'
	fmt.Println(ourRune)

}
