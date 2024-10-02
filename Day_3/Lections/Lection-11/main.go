package main

import (
	"fmt"
)

func main() {

	// ОБЫЧНО так
	// var array [3]int

	// array[0] = 100
	// array[1] = 200
	// array[1] = 300

	// a := 6
	// fmt.Println(a)

	// При создании массива необходмио определить размер на этапе компиляции.
	//var array[a] int - no

	// Слайсы
	// 1. Слайс - это динамическая обвязка над массивом

	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)

	// 2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 25, 30, 70}
	fmt.Println("secondSlice:", secondSlice)

	// 3. Изменение элементов
	originArr := [...]int{30, 40, 50, 60, 70}
	thirdSlice := originArr[1:4]

	for i := range thirdSlice {
		thirdSlice[i]++
	}
	fmt.Println("originArr:", originArr)
	fmt.Println("thirdSlice:", thirdSlice)

	// 4. Один массив и два производных среза

	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Before changing originArr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++

	fmt.Println("After changing originArr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

	// 5. Срез как встроеный тип
	// type slice struct {
	// Lenght int
	// Capacity int
	// ZeroElement *byte
	//}

	// 6. Длинна и емкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "Lenght:", len(wordsSlice), "Capacity", cap(wordsSlice))

	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "Lenght:", len(wordsSlice), "Capacity", cap(wordsSlice))

	// Capacity(cap) - это значение, показывающее сколько элементов можно
	// добавить в слайс, не выделяя доп памяти
	// Если нет места под новый элемент, то выделяется cap * 2

	// numerics := []int{1, 2}

	// for i := 0; i < 200; i++ {
	// 	if i%5 == 0 {
	// 		fmt.Println("Current len:", len(numerics), "Cap:", cap(numerics))
	// 	}
	// 	numerics = append(numerics, i)
	// }

	// Важно! После выделения памяти под новый массив, ссылки со старого будут
	// перенесены в новый

	numArr := []int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 0
	fmt.Println(numArr)
	fmt.Println(numSlice)

	// 7. Как создавать слайсы наиболее эффективно
	// make() - это функция позволяет более детально создать срезы

	sl := make([]int, 10, 15)

	//[]int - тип коллекции
	// 10 длина
	// 15 емкость
	// Сначала иниц массив arr = [15]int
	// затем по нему делается срез arr[0:10]
	// После этого он возвращается
	fmt.Println(sl)

	// 8. Добавление элементов в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherWords := []string{"four", "five", "six"}
	myWords = append(myWords, anotherWords...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println(myWords)

	// 9. Многомерный срез
	mySlice := [][]int{
		{1, 2},
		{2, 3, 4, 5},
		{10, 20, 30},
		{},
	}
	fmt.Printf("Values: %v, as is: %#v", mySlice, mySlice)

	// 10. Слайсы рун
	runeHexSlice := []rune{0x45,0x46,0x47,0x48}
	myStrFromRune := string(runeHexSlice)
	fmt.Println("Runes:",runeHexSlice, "Slice:", myStrFromRune)

	// 10.1 Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 't'}
	myStrFromLiteralRunes := string(runeLiteralSlice)
	fmt.Println("Runes: ",runeLiteralSlice, "Slice: ", myStrFromLiteralRunes)


	// 10.4 Синсахар
	myDecimalByteSlice :=[]byte{100,101,102,103}
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

}
