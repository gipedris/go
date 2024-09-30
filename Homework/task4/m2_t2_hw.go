/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Println("Введите четырезначное число: ")
	fmt.Scan(&input)

	a1 := input / 1000 % 10
	a2 := input / 100 % 10
	a3 := input / 10 % 10
	a4 := input % 10

	if a1 == a4 && a2 == a3 {
		fmt.Println(input, " - палиндром")
		return
	}

	fmt.Println(input, " - не палиндром")
}
