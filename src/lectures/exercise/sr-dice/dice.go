//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func main() {
	var numDice int = 2
	var numRolls int = 10
	var numSides int = 6

	for i := 1; i <= numRolls; i++ {
		rollSum := 0
		fmt.Print("Roll #", i, ": ")
		for j := 1; j <= numDice; j++ {
			rollResult := roll(numSides)
			if j > 1 {
				fmt.Print(" + ", rollResult)
			} else {
				fmt.Print(rollResult)
			}

			rollSum += rollResult
		}
		fmt.Println(" =", rollSum)

		switch {
		case rollSum == 2 && numDice == 2:
			fmt.Println("Snake Eyes")
		case rollSum == 7:
			fmt.Println("Lucky 7")
		case rollSum%2 == 0:
			fmt.Println("Even")
		default:
			fmt.Println("Odd")
		}

	}

}
