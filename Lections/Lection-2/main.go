package main

import "fmt"

func main() {
	var age int
	fmt.Println("My age is:", age)

	var height int = 138
	fmt.Println("My height is:", height)

	var uid = 12345
	fmt.Println("My uid is:", uid)
	fmt.Printf("%T\n", uid)

	// constants
	const price, tax float32 = 275.00, 27.50 // Типизированные
	const quantity, inStock = 2, true // Нетипизированные
	fmt.Println("Total:", 2 * quantity * (price + tax))

	// Vars
	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println(cost + total)

	// Error block
	var value = 275.00
	var taxes float32 = 27.50
	fmt.Println( value + float64(taxes)) 

	// Короткое присваивание
	value, new_value := 3.12, 121
	fmt.Println("Result:", new_value )

	// Ввод данных
	var (
		number int
		s string
	)

	fmt.Scan(&number)
	fmt.Scan(&s)
	fmt.Println(number, s)

}
