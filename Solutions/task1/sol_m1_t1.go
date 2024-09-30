// Задача №1
// Программа запрашивает имя пользователя и возраст
// Нужно вывести имя и возраст за вычетом одного года

package main

import "fmt"

func main() {

	// Ввод данных
	var (
		age  int
		name string
	)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	fmt.Println("Enter your name:")
	fmt.Scan(&name)
	fmt.Println("Your age is:", age - 1)
	fmt.Println("Your name is:", name)
	fmt.Printf("Your name is %s.\nYour age is %d.\n", name, age - 1)
}
