//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

// Product
//  - Products must include the price and the name
type Product struct {
	name  string
	price float32
}

// getInfo - print the last item and list info
func getInfo(list [4]Product) {
	var totalItems int
	var totalCost float32
	for i := 0; i < len(list); i++ {
		item := list[i]
		if item.name != "" && item.price > 0 {
			totalItems++
			totalCost += item.price
		}
	}

	fmt.Println("Last item:", list[totalItems-1])
	fmt.Printf("%d items, totalling %.2f USD\n", totalItems, totalCost)
}

func main() {
	//* Using an array, create a shopping list with enough room
	//  for 4 products
	//* Insert 3 products into the array
	var shoppingList [4]Product = [4]Product{
		{"Apple", 1.99},
		{"Bananas", 4.95},
		{"Milk", 0.58},
	}
	//* Print to the terminal:
	//  - The last item on the list
	//  - The total number of items
	//  - The total cost of the items
	getInfo(shoppingList)

	//* Add a fourth product to the list
	shoppingList[3] = Product{"Watermelon", 4.00}
	//  and print out the information again
	getInfo(shoppingList)
}
