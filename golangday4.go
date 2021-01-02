package main

import "fmt"

func main() {

	var number [10][10]int
	fmt.Print("enter the number of rows : ")
	var row int
	fmt.Scanln(&row)
	fmt.Print("enter the number of columns : ")
	var colm int
	fmt.Scanln(&colm)
	fmt.Println("enter the numbers")
	for i := 0; i < row; i++ {
		for j := 0; j < colm; j++ {
			fmt.Scanf("%d", &number[i][j])
		}
	}
	fmt.Println("matrix is \n")
	for i := 0; i < row; i++ {
		for j := 0; j < colm; j++ {
			fmt.Printf("%d", number[i][j])
		}
		fmt.Println("\n")
	}

}
