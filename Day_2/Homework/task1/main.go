// 3.1 Пользователь вводит числа a и b (b больше a).
//     Вывести все целые числа, расположенные между ними

package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите число a")
	fmt.Scan(&a)
	fmt.Println("Введите число b")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Ошибка: a > b")
		return
	}

	fmt.Println("Все целые числа, расположенные между a и b:")

	for a += 1; a < b; a++ {
		fmt.Println(a)
	}
	
}


