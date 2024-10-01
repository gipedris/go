package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// Numerics. Integers
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64

	var a int = 32
	b := 92

	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a + b)

	//Тип данных через форматирование
	fmt.Printf("type of a: %T and type of b: %T\n", a, b)

	// Узнаем сколько байт в памяти занимает
	// int - платформозависимый
	fmt.Printf("type %T size %d bytes\n", a, unsafe.Sizeof(a))
	fmt.Printf("type %T size %d bytes\n", b, unsafe.Sizeof(b))

	//Если выполняются арифметические операции над int и intX
	// то используйте операции привидения
	
	var c64 int64 = 16_213_897
	var usualInt int = 123_456
	fmt.Println(c64 + int64(usualInt))

	// Аналогичным обзаом устроены uint8, uint16, uint32, uint64

	// Numerics. Float
	// float32, float64

	floatFirst, floatSecond := 5.67, 12.54
	fmt.Printf("type of a: %T and type of b: %T\n", floatFirst, floatSecond)
	fmt.Printf("type %T size %d bytes\n", floatFirst, unsafe.Sizeof(floatSecond))
	fmt.Printf("type %T size %d bytes\n", floatFirst, unsafe.Sizeof(floatSecond))
	sum := floatFirst + floatSecond
	sub := floatFirst + floatSecond
	fmt.Printf("Sum: %3.f Sub: %3.f\n", sum, sub ) 
	fmt.Println(0.1 * 3 == 0.3)
	fmt.Println(math.Pow(floatFirst, floatSecond))



	//Strings
	name := "Федя"
	lastname := "Pupkin"
	fi := name + " " + lastname
	fmt.Println(fi)
	//кол-во байт
	fmt.Println("lenght", len(name))
	fmt.Println("lenght", len(lastname))

	//Rune - руна. Это один utf-ный символ
	fmt.Println("кол-во символов", utf8.RuneCountInString(name))
	fmt.Println("кол-во символов", utf8.RuneCountInString(lastname))

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "asd"
	fmt.Println(strings.Contains(totalString, subString))

	// rune - это alias int32
	var sampleRune rune
	var anotherRune = 'Q' //для иниц используйте кавычки
	var thirdRune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char: %c\n", sampleRune)
	fmt.Println(anotherRune)
	fmt.Printf("Rune as char: %c\n", anotherRune)
	fmt.Println(thirdRune)
	fmt.Printf("Rune as char: %c\n", thirdRune)

	// -1 if first < second, 0 if ==, 1 if first > second
	fmt.Println(strings.Compare("abcd", string('a')))

	var aByte byte // alias uint8
	fmt.Println("Byte", aByte)

}
