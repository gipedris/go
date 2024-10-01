// Задача №2
// Вход: трехзначное число.
// Найти первую, вторую и последнюю цифры заданного трехзначного числа.

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
