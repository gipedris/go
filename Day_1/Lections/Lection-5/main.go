package main

import (
	"fmt"
	"strings"
)

func main() {
	//Boolean => default: false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// Boolean operands

	aBoolean, bBoolean := true, true
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT", !aBoolean)

	//Классический условный оператор
	/*
		if condition {
			//body
		}
	*/

	//Классический условный оператор блоком else
	/*
		if condition {
			//body 1
		}else {
			//body 2
		}
	*/

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Number", value, "is even.")
	} else {
		fmt.Println("Number", value, "is odd.")
	}

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	// init in if
	// присваивание только :=
	// переменная видна только внутри блока (if, else if или else)
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	//Не идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
	} else {
		fmt.Println("width <= 100")
	}

	//Идеоматично
	if width := 100; width > 100 {
		fmt.Println("width > 100")
		return
	}

	fmt.Println("width <= 100")

}
