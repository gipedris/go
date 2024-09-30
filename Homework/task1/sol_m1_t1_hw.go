/*
Задача №1
Вход:

	расстояние(50 - 10000 км),
	расход в литрах (5-25 литров) на 100 км и
	стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import (
	"fmt"
)

func main() {

	const gas_price int16 = 48

	var (
		distance        float64
		consumption     float64
		result          float64
		distanceHint    string = "Укажите расстояние в диапазоне от 50 до 10000 км"
		consumptionHint string = "Укажите расход в литрах в диапазоне от 5 до 25 литров на 100 км"
		error           string = "Неправильные данные:"
	)

	fmt.Println(distanceHint)
	fmt.Scan(&distance)
	if distance < 50 || distance > 10000 {
		fmt.Println(error, distanceHint)
		return
	}

	fmt.Println(consumptionHint)
	fmt.Scan(&consumption)
	if consumption < 5 || consumption > 25 {
		fmt.Println(error, consumptionHint)
		return
	}

	result = (distance / 100) * consumption * float64(gas_price)
	fmt.Printf("Cтоимость поездки: %.2f ₽\n", result)
}
