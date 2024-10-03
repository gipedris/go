package main

import (
	"fmt"
)

func main() {
	// 1. Анонимная функция
	anon := func(a, b int) int {
		return a + b

	}
	fmt.Println(anon(10, 20))
	// fmt.Println("", bigFunction(10, 20))

	var command string = "substraction"
	res := calcAndReturnValidFunc(command)
	fmt.Println("", command, res(10, 20))

}

// 2. Анонимная функция внутри явной
// func bigFunction(aArg, bArg, int) int {
// 	return func(a, b int) int {return a + b + 100}(aArg, bArg)
// }

// 3.

func calcAndReturnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substraction" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}

}

// 5. Функция как входящий параметр в другой функции

// func receiveFuncAndReturnValue(f func(a, b int) int) int {
// 	var intVarA, intVarB int
// 	intVarA = 100
// 	intVarB = 200

// }

