package main

import "fmt"

func main() {

	fmt.Println("enter first number : ")
	var x int
	fmt.Scanf("%d\n", &x)

	fmt.Println("enter second number : ")
	var y int
	fmt.Scanf("%d\n", &y)

	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y && x == y)
	fmt.Println(!(x <= y))
	fmt.Println(x == y || x < y)

}
