package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(num int64, i int, value int) int64 {
	if value == 1 {
		return num | (1 << i)
	} else if value == 0 {
		return num &^ (1 << i)
	}
	return num
}

func main() {
	var number int64
	number = 42
	bitIndex := 2
	bitValue := 1

	fmt.Printf("Исходное число в двоичной системе: %b\n", number)

	number = setBit(number, bitIndex, bitValue)

	fmt.Printf("Число после установки бита: %b\n", number)
}
