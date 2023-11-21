package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func stringUnique(s string) bool {
	temp := make(map[rune]struct{})

	for _, letter := range strings.ToLower(s) {
		if _, ok := temp[letter]; ok {
			return false
		}
		temp[letter] = struct{}{}
	}

	return true
}

func main() {
	example1 := "abcd"
	fmt.Println(example1, "is", stringUnique(example1))

	example2 := "abCdefAaf"
	fmt.Println(example2, "is", stringUnique(example2))

	example3 := "aabcd"
	fmt.Println(example3, "is", stringUnique(example3))
}
