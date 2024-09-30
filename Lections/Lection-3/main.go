package main

import (
	"fmt"
	"math"
)

func main() {
	var result float64
	var d, e float64 = 10, 4
	var check float64 = 20 / 3

	a := 42
	b := 20
	c := a / b
	result = d / e

	fmt.Printf("type of c %T and value of c %v\n", c, c)
	fmt.Printf("type of result %T and value of result %v\n", result, result)
	fmt.Printf("type of check %T and value of check %v\n", check, check)

	fmt.Println("Increment: ")
	// check = check + 1
	check++
	fmt.Println("check = ", check)
	fmt.Println("Decrement: ")
	//c = c -1
	c--
	fmt.Println("c =", c)

	// Операция взяти остатка
	newCheck := int(check) % 3
	fmt.Println("newCheck:", newCheck)

	// Операция возведение в степень из модуля math
	total := math.Pow(2, 10)
	fmt.Printf("Total = %.1f\n", total)

}
