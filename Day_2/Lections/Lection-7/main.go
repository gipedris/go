package main

import (
	"fmt"
	"strings"
)

func main() {
	
	// for init; condition; post {
	// init блок инициализации переменных
	// condition - условие (если условие верно - тело цикла выполняется
	// если нет - то цикл завершает работу)
	// post - изменение переменной цикла (инкремент или декремент)
	// }

	for i := 0; i <= 5; i++ {
		fmt.Println("Current value of i:", i)
	}

	// важно в качестве init можно использовать только через :=

	// break - прерывает текущее выполнение цикла и предает управление
	// следующим за циклом инструкциям

	for i := 11; i <= 22; i++ {
		if i > 20 {
			break
		}
		fmt.Println("Current value of i:", i)
	}

	fmt.Println("After loop with BREAK")

	// continue - прерывает текущее выполнение цикла и предает управление
	// следующей итерации цикла

	for i := 105; i <= 120; i++ {
		if i > 110 && i <= 115 {
			continue
		}
		fmt.Println("Current value of i:", i)
	}

	fmt.Println("After loop with CONTINUE")

	// Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println("Выше должен быть треугольник")

	// бывают ситуации когда нужно прервать сразу оба цикла
	// здесь помогут лейблы. Лейблы - син сахар

outer:
	for i := 0; i < 2; i++ {
		for j := 1; j <= 3; j++{
			fmt.Printf("i:%d and j: %d and sum i+j:%d\n", i, j, i + j)
			if i == j {
				break outer // остановить все циклы
			}
		}
	}
	fmt.Println("After outer stop")

	// шаблон цикла for как while (модификация for)
	// Классический цикл while do

	var loopVar int = 0
	//while (loopVar < 10){
	// body
	// loopVar++
	//}
	for loopVar < 10 {
		fmt.Println("In while like loop", loopVar)
		loopVar++
	}

	// 2. Классический бесконечный цикл
	var password string
outer2:
	for{
		fmt.Println("Insert password")
		fmt.Scan(&password)
		if strings.Contains(password, "1234"){
			fmt.Println("Bad pass. Try again")
		} else {
			fmt.Println("Password accepted")
			break outer2
		}
	}

	// 3. Цикл с множественными переменными
	for x, y := 0, 1; x < 5 && y < 8; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

}
