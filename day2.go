package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Enter the value of x:")
	var x float64
	fmt.Scanf("%f\n", &x)

	fmt.Println("Enter the value of y:")
	var y float64
	fmt.Scanf("%f\n", &y)

	fmt.Println("Enter the value of z:")
	var z float64
	fmt.Scanf("%f\n", &z)

	fmt.Println("Enter the value of aa:")
	var aa float64
	fmt.Scanf("%f\n", &aa)

	var result float64
	result = ((math.Sqrt(x) + math.Sqrt(y)) * z) / aa
	fmt.Println("Result is ", result)
}
