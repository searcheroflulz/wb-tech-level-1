package main

import "fmt"

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func reverseString(s string) string {
	runes := []rune(s)
	len := len(runes) - 1

	for i, j := 0, len; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	inputString := "tenet"

	fmt.Printf("Исходная строка: %s\n", inputString)

	reversedString := reverseString(inputString)

	fmt.Printf("Перевернутая строка: %s\n", reversedString)
}
