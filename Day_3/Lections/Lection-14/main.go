package main

// 1. Явная функция - локальный блок когда, имеющий имя
// Функцию необходимо определить + вызвать
// 2. Сигнатура функции
// func functionName(params type) typeReturnValue{
// function body
//}

import (
	"fmt"
)

func main() {
	// 4. Вызов функции
	res := add(10, 20)
	fmt.Println(res)
	fmt.Println(mult(1, 2, 3, 4))
	per, square := rectangleParameters(12, 4)
	fmt.Println(per, square)
	per, area := namedReturn(12, 4)
	fmt.Println(per, area)
	fmt.Println(funcWithReturn(1, 2))

	emptyReturn(12)
	helloVariadic(1)
	helloVariadic(1, 2)

	fmt.Println("sum variadic:", sumVariadic(1, 2))
	sliceNumbers := []int{5, 6, 7, 8}
	fmt.Println("sum slice variadic:", funcSlice(sliceNumbers))

	resSum2 := sumVariadic(sliceNumbers...)
	fmt.Println("sum slice variadic:", resSum2)

}

// 3. Простейшая функция (определить можно как до момента ее вызова в main),
// так и в любом месте пакета, главное чтобы она была определена в принципе

// add function add two numbers
func add(a int, b int) int {
	result := a + b
	return result
}

// 5. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	return a * b * c * d
}

// 6. Возврат больше одного значения
func rectangleParameters(a, b int) (int32, int) {
	perimeter := int32(2 * (a + b))
	area := a * b
	return perimeter, area
}

// 7. Именнованый возврат значений
func namedReturn(widht, height int) (perimeter int32, area int) {
	perimeter = int32(2 * (widht + height))
	area = widht * height
	return
}

// 8. При вызове оператора return функция прекращает свое выполнение

func funcWithReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}
	if a == b {
		return a, true
	}
	return 0, false
}

// 9. Что вернет функция в случае, если return не указан или он пустой
func emptyReturn(a int) {
	fmt.Println("i am f with params", a)
}

// 10. Variadic params (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("%v\n", a)
}

// 11. Смешивание параметров с variadic
// Континуальный параметр всегда самый последний
// Variadic parameter один на всю функцию

func someStrings(a, b int, words ...string) {
	fmt.Printf("%v\n", a)
}

// 12. Передача слайса с variadic

func sumVariadic(nums ...int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}

func funcSlice(nums []int) int {
	var sum int
	for _, number := range nums {
		sum += number
	}
	return sum
}
