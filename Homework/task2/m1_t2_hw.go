/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import (
	"fmt"
)

func main() {

	var input int
	fmt.Println("Введите трехзначное число: ")
	fmt.Scan(&input)
	first := input / 100
	second := input % 100 / 10
	third := input % 10
	fmt.Printf("%v%v%v\n", third, second, first)
}
