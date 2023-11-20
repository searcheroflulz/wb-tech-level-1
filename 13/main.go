package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 5
	b := 10
	fmt.Printf("A = %v, B = %v\n", a, b)
	a, b = b, a
	fmt.Printf("A = %v, B = %v\n", a, b)
}
