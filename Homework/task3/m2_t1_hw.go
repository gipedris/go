/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1 9 2
Выход: 1 2 9

Про sort мы пока ничего не знаем!
Это задача на условный оператор
*/
package main

import (
	"fmt"
)

func main() {

	var (
		a, b, c, x int
	)

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)
	fmt.Println("Введите третье число: ")
	fmt.Scan(&c)
	fmt.Printf("Ввод: %v %v %v\n", a, b, c)

	if b > c {
		x = b
		b = c
		c = x
	}

	if a > c {
		x = a
		a = c
		c = x
	}

	if a > b {
		x = a
		a = b
		b = x
	}

	fmt.Printf("Вывод: %v %v %v\n", a, b, c)
}

/*
	9 6 3
	9 3 6
	6 3 9
	3 6 9
*/
