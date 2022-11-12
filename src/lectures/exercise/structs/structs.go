//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

// Rectangle
//* Create a rectangle structure containing its coordinates
type Rectangle struct {
	x0, y0, x1, y1 float64
	area           float64
	perimeter      float64
}

// calculateArea
//* Using functions, calculate the area and perimeter of a rectangle,
//  - The functions must use the rectangle structure as the function parameter
func calculateArea(r Rectangle) float64 {
	return math.Abs(r.x1-r.x0) * math.Abs(r.y1-r.y0)
}

func calculatePerimeter(r Rectangle) float64 {
	return 2 * (math.Abs(r.x1-r.x0) + math.Abs(r.y1-r.y0))
}

func printRect(r Rectangle) {
	fmt.Println("Area is:", r.area)
	fmt.Println("Perimeter is:", r.perimeter)
}

func main() {
	rect := Rectangle{x0: 0.0, x1: 4.0, y0: 0.0, y1: 8.0}
	rect.area = calculateArea(rect)
	rect.perimeter = calculatePerimeter(rect)
	//  - Print the results to the terminal
	printRect(rect)
	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.x1 = rect.x1 * 2
	rect.y1 = rect.y1 * 2
	rect.area = calculateArea(rect)
	rect.perimeter = calculatePerimeter(rect)
	//  - Print the new results to the terminal
	printRect(rect)
}
