package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func getSquare(wg *sync.WaitGroup, num int, ch chan int) {
	defer wg.Done()
	square := math.Pow(float64(num), 2)
	ch <- int(square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go getSquare(&wg, num, resultChan)
	}
	wg.Wait()
	close(resultChan)

	var result int
	for square := range resultChan {
		result += square
	}
	fmt.Printf("Sum of squares = %v", result)
}
