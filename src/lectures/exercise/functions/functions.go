//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("Hello", name)
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func returnMessage() string {
	return "Hello from returnMessage"
}
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(x,y,z int) int {
	return x+y+z
}
//* Write a function that returns any number
func returnAnyNumber() int {
	return 5
}
//* Write a function that returns any two numbers
func returnAnyTwoNumbers() (int, int) {
	return 5, 6
}
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

func main() {
	greet("Serzhan")
	fmt.Println(returnMessage())
	sum := add(1, 2, 4)
	fmt.Println("Sum is", sum)
	a, b := returnAnyTwoNumbers()
	sum = add(returnAnyNumber(), a, b)
	fmt.Println("Sum again", sum)

}
