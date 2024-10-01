/*
Задача №1.
Программа получает на вход последовательность из 5 целых чисел.

Вам нужно определить вид последовательности:
 - возрастающая
 - убывающая
 - случайная
 - постоянная

В качестве ответа следуют выдать прописными латинскими буквами тип последовательности:
1. ASCENDING (строго возрастающая)
2. WEAKLY ASCENDING (нестрого возрастающая, то есть неубывающая)
3. DESCENDING (строго убывающая)
4. WEAKLY DESCENDING (нестрого убывающая, то есть невозрастающая)
5. CONSTANT (постоянная)
6. RANDOM (случайная)

Примеры входных и выходных данных:
In: 1 2 4 9 11   Out: ASCENDING 2
In: 11 9 4 2 -1  Out: DESCENDING -2
In: 3 8 8 11 12  Out: WEAKLY ASCENDING 1
In: 2 -1 7 21 1  Out: RANDOM not all
In: 5 5 5 5 5    Out: CONSTANT 0

Подсказка: используем метод строки strings.Split()
*/

package main

import (
	"fmt"
)

func main() {

	// var arr [5]int
	// var status int

	fmt.Println("Введите 5 целых чисел:")
	// for i := 0; i < 5; i++ {
	// 	fmt.Scan(&arr[i])
	// 	if i > 0 {
	// 		if arr[i] > arr[i -1] {
	// 			status += 2
	// 		}
	// 		if arr[i] == arr[i -1] {
	// 			status += 1
	// 		}

	// 		if arr[i] < arr[i -1] {
	// 			status += -2
	// 		}
	// 	}
	// }
	
	// fmt.Println(status)

}
