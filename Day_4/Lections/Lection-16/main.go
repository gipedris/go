package main

import (
	"fmt"
)

func main()  {
	// 1. Указатели - переменная, хранящая адрес другой переменной

	// 2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable // взятие адреса в памяти
	fmt.Printf("Type: %T Value: %v\n", pointer, pointer)

	// 3. А какое нулевое значение для указателя
	var zeroPointer *int // nil
	fmt.Printf("Type: %T Value: %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
	}
	fmt.Printf("Type: %T Value: %v\n", zeroPointer, zeroPointer)

	// 4. Разыменовывание указателя(получение значения)
	// *pointer - возвращает значение

	var numValue int = 32
	pointerToNum := &numValue
	fmt.Printf("Value: %v Value: %v\n", &pointerToNum, *pointerToNum)

	// 5. Создание указателей на нулевые значения типов
	// var zeroVar int
	// var zeroPoint *int = &zeroVar
	// fmt.Println(zeroPoint)

	zeroNewPoint := new(int) 
	fmt.Printf("Value: %v Value: %v\n", &zeroNewPoint, *zeroNewPoint)

	// 6. Изменение значения хранимого по адресу через указатель
	zeroPtoInt := new(int)
	fmt.Printf("Value: %v Value: %v\n", &zeroPtoInt, *zeroPtoInt)
	*zeroPtoInt += 40
	fmt.Printf("Value: %v Value: %v\n", &zeroPtoInt, *zeroPtoInt)

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100

	fmt.Printf("Value: %v\n", b)

	// 7. Указательная арифметка ПОЛНОСТЬЮ ОТСУТСВУЕТ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес ходить по памяти


	// 8. Передача указателей в фунции
	// Большой прирост производительности за сч)т тогоб что перда)тся не значение
	// а адрес в памяти за которым хранится значение переменной
	sample := 1
	fmt.Println("Sample:", sample)
	changeParam(&sample)
	fmt.Println("Sample:", sample)

	// 9. Возврат указателя из функции
	ptr1 := returnPointer()
	ptr2 := returnPointer()

	fmt.Printf("ptr1 Type: %T, addres: %v value %v\n", ptr1,ptr1,*ptr1)
	fmt.Printf("ptr2 Type: %T, addres: %v value %v\n", ptr2,ptr2,*ptr2)
	
	arr := [3]int{1,2,3}
	fmt.Println("Before", arr)
	mutation(&arr)
	fmt.Println("After", arr)
}

// 8.1 Определение функции, принимающей указатель как параметр
func changeParam(value *int) {
	*value += 100
}

// 9.1 

func returnPointer() *int {
	var numeric int = 321
	return &numeric // В момент возврата премещается в кучу
}

// 10. Указатели на массивы

func mutation(arr *[3]int){
	// (*arr)[1] = 900
	arr[1]=900
	arr[2]=1000
}

// 11. 
func mutationSlice(sls []int){
	sls[1] = 100
	sls[2] = 200
}

