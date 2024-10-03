package main

import (
	"fmt"
)

// 1. Структура - заименованный набор полей(состояний),
// определяющий новый тип данных

// 2. Определение структуры явно

type Student struct {
	firstName string
	lastname string
	age int
}

// 11. Создание экземпляра с анонимными полями
type Human struct {
	firstName string
	lastname string
	string
	int
	bool
}

// 3. Если имеется ряд состояний одинакового типа, то можно сделать так

type AnotherStuden struct{
	firstName, lastName, groupName string
	age, courseNumber int
}


func PrintStudent(std Student){
	fmt.Println("================")
	fmt.Println("First Name:", std.firstName )
	fmt.Println("Last Name:", std.lastname )
	fmt.Println("Age:", std.age )
}


func main()  {
	// 4. Создание экземпляра
	student := Student{
		firstName: "Fedya",
		lastname: "Petrov",
		age: 21,
	}
	PrintStudent(student)

	student2 := Student{"Petr", "Ivanov", 19} //порядок указания свойств как в структуре
	PrintStudent(student2)

	// 5. Что если не все поля определены

	student3 := Student{
		firstName: "Vasya",
	}

	PrintStudent(student3)

	// 6. Анонимные структуры
	anonStudent := struct {
		age int
		groupID int
		professorName string
	}{
		age: 23,
		groupID: 2,
		professorName: "Alex",
	}
	fmt.Println(anonStudent)

	// 7. Как организован доступ к полям(состояниям) и их модификация
	studMark := Student{"Mark", "Ivanov", 19}
	fmt.Println("================")
	fmt.Println("First Name:", studMark.firstName )
	fmt.Println("Last Name:", studMark.lastname )
	fmt.Println("Age:", studMark.age )
	studMark.age += 2
	fmt.Println("New Age:", studMark.age )

	// 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student
	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	// 9. Указатели
	studPointer := &Student {
		firstName: "Igor",
		lastname: "Daaad",
		age: 25,
	}
	fmt.Println("Value of studPointer:", studPointer)

	secondPointer := studPointer
	(*secondPointer).age += 20

	fmt.Println("Value of studPointer:", studPointer)

	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	// 10. Работа с доступом к полям через указатель
	fmt.Println("Age via (*...).age:", (*studPointer).age)
	// Неявно происходит разыменование и обращение к полю
	fmt.Println("Age via .age:", studPointer.age)

	// 12. Создание экземпляра с анонимными полями
	human := & Human{
		firstName: "Bob",
		lastname: "Smirnoff",
		string: "Additional info",
		int: -1,
		bool: true,
	}
	fmt.Println(human)
	fmt.Println("Anon field string:", human.string)
}

