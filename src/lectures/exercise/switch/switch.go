//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)
	switch age := gen.Intn(24); {
	case age == 0:
		fmt.Println("newborn")
	case age < 4:
		println("toddler")
	case age < 13:
		println("child")
	case age < 18:
		println("teenager")
	default:
		println("adult")
	}
}
