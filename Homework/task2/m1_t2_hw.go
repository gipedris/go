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
	// сколько целых сотен
	first := input / 100
	fmt.Println("Первая цифра", first)
	// сколько целых десятков
	second := input % 100 / 10
	fmt.Println("Вторая цифра", second)
	// сколько единиц
	third := input % 10
	fmt.Println("Третья цифра", third)
}
