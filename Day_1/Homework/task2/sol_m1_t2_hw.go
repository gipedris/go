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

	var input int16
	fmt.Println("Введите трехзначное число: ")
	fmt.Scan(&input)
	first := input / 100 % 10
	second := input / 10 % 10
	third := input % 10
	fmt.Printf("%v%v%v\n", third, second, first)
}
