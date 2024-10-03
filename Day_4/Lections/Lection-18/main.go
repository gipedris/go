package main

import (
	"fmt"
)

// 1. Вложенные структуры
// Это использование использование одной структуры, как тип поля для другой.

type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// 4. Встроенные структуры

type Professor struct {
	firstName  string
	lastName   string
	age        int
	greatWorks string
	University
}

func main() {

	// 2. Создание экз с вложениями
	student := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "good Un",
			infoLong:  "Its adsrgfafegadfgadfg",
		},
	}
	// 3. Получение доступа к вложенным полям
	fmt.Println("firstName:", student.firstName)
	fmt.Println("lastName:", student.lastName)
	fmt.Println("yearBased:", student.university.yearBased)
	fmt.Println("infoShort:", student.university.infoShort)
	fmt.Println("infoLong:", student.university.infoLong)

	// 5. Создание экз с встроенной структурой
	prof := Professor{
		firstName: "Anatoly",
		lastName: "Smironv",
		age: 35,
		greatWorks: "F5",
		University: University{
			yearBased: 1734,
			infoShort: "re",
			infoLong: "asdgfasdfgadsfga",
			age: 2024 - 1734,	
		},
	}
	
	// 6.
	fmt.Println("firstName:", prof.firstName)
	fmt.Println("lastName:", prof.lastName)
	fmt.Println("yearBased:", prof.yearBased)
	fmt.Println("infoShort:", prof.infoShort)
	fmt.Println("infoLong:", prof.infoLong)
	fmt.Println("age:", prof.University.age)
	
	// 7. Сравнение экземпляров
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)


}
