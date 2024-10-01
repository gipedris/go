package main

import (
	"fmt"
)

func main() {

	// Массивы. Основа
	// 1. Определение массива
	// Создание массива из пяти значений
	//При инициализации массива важно передатьинформацию
	//- сколько элементов в нем будет

	var arr [5]int
	fmt.Println("This is my array", arr)

	// 2. Присваивание значения элементу массива
	// Необходимо обратиться  к элементу массива через его индекс
	// arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = 580
	arr[4] = 1002

	fmt.Println("Array after change", arr)

	// 3. Создание массива с указанием элементов на месте
	newArr := [5]int{10, 20, 30}
	fmt.Println("Short declare init", newArr)

	// 4. Создание массива через инициализацию переменных
	arrWithValue := [...]int{10, 20, 30, 40}
	fmt.Println("Auto init", arrWithValue, len(arrWithValue))
	arrWithValue[0] = 10_000
	fmt.Println("Auto init", arrWithValue, len(arrWithValue))

	//5. Массив это набор значений. При всех манипуляциях массив копируется
	// (жестко на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("1", first)
	fmt.Println("2", second)

	//6. Масси и его размер это 2 составляющие одного типа

	//7. итерирование по массиву
	floatArray := [...]float64{12.5, 10.0, 12.1, 15.2, 13.51}
	for i := 0; i < len(floatArray); i++ {
		fmt.Printf("%d element of array is %.2f:\n", i + 1, floatArray[i])
	}

	//8. Итерирование через range
	fmt.Println("Итерирование через range")

	var sum float64
	for idx, value := range floatArray{
		fmt.Printf("%d element of array is %.2f:\n", idx + 1, floatArray[idx])
		sum += value
	}
	fmt.Println("Total sum:", sum)


	//9. Игнорирование idx при range
	fmt.Println("Игнорирование idx при range")

	var summa float64
	for _, value := range floatArray{
		summa += value
	}
	fmt.Println("Total sum 9:", summa)

	// 10. Многомерные массивы
	words := [2][2]string{
		{"Mark", "alice"},
		{"Viktor", "Imma"},
	}
	fmt.Println("Multidimensional array:", words)

	// 11. Итерирование многомерного массива
	for _, innerArr := range words {
		for _, word := range innerArr{
			fmt.Printf("Word: %s ", word)
		}
		fmt.Println()
	}
}
