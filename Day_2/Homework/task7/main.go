// 3.7 Вводим строку без знаков препинания(5 слов).
//     Найти самое длинное слово в строке и вывести сколько в нем букв.

// Пример:
// In: Скажи как дела в учебе
// Out: учебе, 5

// In: Закрепляем циклы в языке Golang
// Out: Закрепляем, 10

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var sentence [5]string
	var max string

	fmt.Println("Введите строку без знаков препинания (5 слов):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&sentence[i])
		if utf8.RuneCountInString(sentence[i]) >= utf8.RuneCountInString(max) {
			max = sentence[i]
		}
	}
	fmt.Printf("%s, %d\n", max, utf8.RuneCountInString(max))

}
