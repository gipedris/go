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
		deposit     float64
		years       int64
		percent     int64
		result      float64
		depositHint string = "Укажите размер вклада в диапазоне от 100 до 1_000_000"
		yearsHint   string = "Укажите кол-во лет в диапазоне от 1 до 100"
		percentHint string = "Укажите процент по вкладу в диапазоне от 1 до 20"
		error       string = "Неправильные данные:"
	)

	fmt.Println(depositHint)
	fmt.Scan(&deposit)
	if deposit < 100 || deposit > 1000000 {
		fmt.Println(error, depositHint)
		return
	}

	fmt.Println(yearsHint)
	fmt.Scan(&years)
	if years < 1 || years > 100 {
		fmt.Println(error, yearsHint)
		return
	}

	fmt.Println(percentHint)
	fmt.Scan(&percent)
	if percent < 1 || percent > 20 {
		fmt.Println(error, percentHint)
		return
	}

	result = deposit * (math.Pow((1 + float64(percent)/100), float64(years)))
	fmt.Printf("Размер вклада с капитализацией: %.2f\n", result)
}
