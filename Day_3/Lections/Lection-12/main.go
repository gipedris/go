package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var name string

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		// Захват в буфер
		name = input.Text()
		// Text() возвращяет элементы помещенные в буфер в виде строки
	}
	fmt.Println(name)

	// Вариант с циклом

	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println(result)
		}
	}

	// Преобразование строкового литерала к чему-нибудь числовому

	numStr := "10"
	nint, _ := strconv.Atoi(numStr)
	fmt.Println(nint)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Println(numInt64)

	// ParseFloat

	// Пример bufio.NewReader
	fmt.Println("Пример bufio.NewReader")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Fields(line)
	fmt.Printf("data value: %v an as is: %#v\n", data, data)
}
