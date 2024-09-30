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
		deposit      float64
		years        int64
		percent      int64
		result       float64
		deposit_hint string = "Укажите размер вклада в диапазоне от 100 до 1_000_000"
		years_hint   string = "Укажите кол-во лет в диапазоне от 1 до 100"
		percent_hint string = "Укажите процент по вкладу в диапазоне от 1 до 20"
		error_hint   string = "Неправильные данные:"
	)

	fmt.Println(deposit_hint)
	fmt.Scan(&deposit)
	if deposit < 100 || deposit > 1000000 {
		fmt.Println(error_hint, deposit_hint)
		return
	}

	fmt.Println(years_hint)
	fmt.Scan(&years)
	if years < 1 || years > 100 {
		fmt.Println(error_hint, years_hint)
		return
	}

	fmt.Println(percent_hint)
	fmt.Scan(&percent)
	if percent < 1 || percent > 20 {
		fmt.Println(error_hint, percent_hint)
		return
	}

	result = deposit * (math.Pow((1 + float64(percent)/100), float64(years)))
	fmt.Printf("Размер вклада с капитализацией: %.2f\n", result)
}
