// 3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)

package main

import "fmt"

func main() {
	var a int

	fmt.Println("Введите число:")
	fmt.Scan(&a)

	fmt.Printf("Таблица умножения на %d:\n", a)

	for i := 1; i < 11; i++ {
		sum := a * i
		fmt.Printf( "%d x %d = %d\n", a, i, sum)
	}
}
