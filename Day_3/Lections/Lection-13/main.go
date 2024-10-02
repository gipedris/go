package main

import (
	"fmt"
	"strings"
)

func main() {

	// 1. Map - это набор ключ:значение
	// Инициализация пустой мапы
	mapper := make(map[string]int)

	fmt.Println(mapper)
	// 2. Добавление элементов в мапу
	mapper["Petr"] = 23
	mapper["Elena"] = 12
	fmt.Println(mapper)

	// 3. Инициализация мапы с указанием пар
	newMapper := map[string]int{
		"Alice": 192,
		"Bob": 134,
	}
	newMapper["Nick"] = 87
	fmt.Println(newMapper)

	// 4. Что может быть ключом в мапе
	// Важно мапа в Гошке не упорядоченная
	// ключом может быть только сравнимый тип (==, !=)

	// 5. Нулево значение мапы
	// var mapZeroValue map[string]int // == nil

	// 6. Получение значений мапы
	testPerson := "Alice"
	fmt.Println("Salary of Alice:", newMapper[testPerson])

	testPerson = "Vasya"
	fmt.Println("Salary of Alice:", newMapper[testPerson])

	// 7. Проверка наличия ключа

	employee := map[string]int{
		"Denis": 0,
		"Alice": 0,
		"Vasya":0,
	}

	//При обращении к ключу получаем два значения

	if value, ok := employee["Denis"]; ok {
		fmt.Println("Denis exists and it's value:", value )
	} else {
		fmt.Println("Denis dosn't exists in map")
	}

	if value, ok := employee["Bob"]; ok {
		fmt.Println("Bob exists and it's value:", value )
	} else {
		fmt.Println("Bob dosn't exists in map")
	}

	// 8. Перебор значений мапы

	fmt.Println(strings.Repeat("=", 15))

	for key, value := range employee {
		fmt.Printf("key: %s and it's value: %d\n", key, value)
	}

	// 9. Как удалять пары
	// 9.1 Удаление сущ пары
	fmt.Println("Before deleting", employee)
	delete(employee, "Vasya")
	fmt.Println("After deleting", employee)

	// 9.2 Удаление несуществующей пары

	if _, ok := employee["Anna"]; ok {
		delete(employee, "Anna")
	}
	
	fmt.Println(employee)

	// 10. Количество пар == длина мапы
	fmt.Println("Pairs amount in map:", len(employee))

	// 11. Мапа - ссылочный тип
	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newWords := words
	newWords["Three"]="Три"
	delete(newWords, "One")
	fmt.Println(newWords)

	// 12. Сравнение мап
	// 12.1 сравнение массивов ()
	if [3]int{1,2,3} == [3]int{1,2,3} {
		fmt.Println("Equal")
	}

	// 12.2 Сравнение слайсов, из-за того что слайс ссылочный тип его можно сравнивать только с nil
	//[]int{1,2,3} == 

	// 12.3 Сравнение мап
	// из-за того что мап ссылочный тип его можно сравнивать только с nil

	aMap := map[string]int{
		"a": 1,
	}


	if aMap == nil {
		fmt.Println("nil")
	}

	// Если слайс или мапа являются составляющими какой-либо структуры
	// то структура несравнима

	
}
