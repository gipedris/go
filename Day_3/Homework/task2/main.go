/*
	Задача №2
	
	Вход:
	Пользователь должен ввести "правильный пароль", состоящий из:
	цифр, букв латинского алфавита(строчные и прописные) и 
	специальных символов  special = "_!@#$%^&"

	Всего 4 набора различных символов.
	В пароле обязательно должен быть хотя бы один символ из каждого набора.
	Длина пароля от 8(мин) до 15(макс) символов.
	Максимальное количество попыток ввода неправильного пароля - 5.
	Каждый раз выводим номер попытки.
	выводить пояснение, почему пароль не принят и что нужно исправить.

	digits = "0123456789"
	lowercase = "abcdefghiklmnopqrstvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special = "_!@#$%^&"

	Выход:
	Написать, что пароль правильный и он принят.

	Пример: 
	хороший пароль -> o58anuahaunH!
	хороший пароль -> aaaAAA111!!!
	плохой пароль -> saucacAusacu8 

	Реализацию оформить через функцию:
	checkPassword(pass string) (bool, errors <- на усмотрение)

*/
package main

import (
	"strings"
	"errors"
	"fmt"
	"unicode"
)

const (
	digits    = "0123456789"
	lowercase = "abcdefghiklmnopqrstvxyz"
	uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
	special   = "_!@#$%^&"
)

func checkPassword(pass string) (bool, error) {
	if len(pass) < 8 || len(pass) > 15 {
		return false, errors.New("длина пароля должна быть от 8 до 15 символов")
	}

	hasDigit := false
	hasLower := false
	hasUpper := false
	hasSpecial := false

	for _, char := range pass {
		switch {
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case isSpecial(char):
			hasSpecial = true
		}
	}

	if !hasDigit {
		return false, errors.New("пароль должен содержать хотя бы одну цифру")
	}
	if !hasLower {
		return false, errors.New("пароль должен содержать хотя бы одну строчную букву")
	}
	if !hasUpper {
		return false, errors.New("пароль должен содержать хотя бы одну прописную букву")
	}
	if !hasSpecial {
		return false, errors.New("пароль должен содержать хотя бы один специальный символ")
	}

	return true, nil
}

func isSpecial(char rune) bool {
	return strings.ContainsRune(special, char)
}

func main() {
	var password string
	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Printf("Попытка %d: введите пароль: ", attempts)
		fmt.Scan(&password)

		valid, err := checkPassword(password)
		if valid {
			fmt.Println("Пароль правильный и он принят.")
			return
		} else {
			fmt.Println("Ошибка:", err)
			if attempts == 5 {
				fmt.Println("Превышено максимальное количество попыток.")
			}
		}
	}
}