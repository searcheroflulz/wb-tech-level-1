package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

/*
Разработать программу, которая перемножает,
делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

func Sum(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

func main() {
	a := big.NewInt(int64(math.Pow(2, 20)) + int64(rand.Intn(5000)))
	b := big.NewInt(int64(math.Pow(2, 20)) + int64(rand.Intn(1000)))

	fmt.Printf("a равно: %v\n", a)
	fmt.Printf("b равно: %v\n", b)

	fmt.Printf("Результат сложения: %v\n", Sum(a, b))

	fmt.Printf("Результат вычитания: %v\n", Sub(a, b))

	fmt.Printf("Результат умножения: %v\n", Mul(a, b))

	fmt.Printf("Результат деления: %v\n", Div(a, b))
}
