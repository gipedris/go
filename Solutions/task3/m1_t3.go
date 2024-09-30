/*
Задача №3

На входе размер вклада(float64), кол-во лет(int64) и процент по вкладу(int64)
Проверить условия (от и до включительно):
вклад от 100 до 1_000_000
кол-во лет от 1 до 100
процент от 1 до 20

и посчитать размер вклада при выполнении условий.
В противном случае вывести сообщение о неправильных данных

Использовать ежегодную капитализацию.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		deposit float64
		years   int64
		percent int64
		result  float64
	)

	fmt.Println("Укажите размер вклада в диапазоне от 100 до 1_000_000")
	fmt.Scan(&deposit)
	fmt.Println("Укажите кол-во лет в диапазоне от 1 до 100")
	fmt.Scan(&years)
	fmt.Println("Укажите процент по вкладу в диапазоне от 1 до 20")
	fmt.Scan(&percent)

	if deposit < 100 || deposit > 1000000 {
		fmt.Println("Неправлиьные данные: вклад от 100 до 1_000_000")
		return
	}

	if years < 1 || years > 100 {
		fmt.Println("Неправлиьные данные: кол-во лет от 1 до 100")
		return
	}

	if percent < 1 || percent > 20 {
		fmt.Println("Неправлиьные данные: кол-во лет от 1 до 100")
		return
	}


	result = deposit * (math.Pow(float64((1 + percent/100)), float64(years)))

	fmt.Println("Размер вклада: ", result)
}
