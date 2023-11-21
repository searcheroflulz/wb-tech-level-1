package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func reverseSentence(s string) string {
	words := strings.Split(s, " ")
	var result string

	for i := len(words) - 1; i >= 0; i-- {
		result += words[i] + " "
	}
	return result
}

func main() {
	inputString := "sun dog snow"
	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reverseSentence(inputString))
}
