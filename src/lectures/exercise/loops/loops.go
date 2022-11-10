//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	//* Print integers 1 to 50, except:
	for i := 1; i <= 105; i++ {
		var w string
		if i%3 == 0 {
			w += "Fizz"
		}
		if i%5 == 0 {
			w += "Buzz"
		}
		if i%7 == 0 {
			w += "Fuzz"
		}
		if len(w) > 0 {
			fmt.Println(w)
		} else {
			fmt.Println(i)
		}
	}
}
