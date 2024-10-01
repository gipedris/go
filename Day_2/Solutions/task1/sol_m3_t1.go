/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for

Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

import (
	"fmt"
)

func main() {
	var number, sum int

	fmt.Println("Введите натуральное число:")
	fmt.Scan(&number)

	for number > 0 {
		sum += number % 10
		number /= 10
	}

	fmt.Println(sum)
}
