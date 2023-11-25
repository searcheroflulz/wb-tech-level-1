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

// считаем квадрат числа и записываем в канал
func getSquare(wg *sync.WaitGroup, num int, ch chan int) {
	defer wg.Done()
	square := math.Pow(float64(num), 2)
	ch <- int(square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	//создаем канал в который будем записывать результаты вычислений (буфер размером с слайс чисел)
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go getSquare(&wg, num, resultChan)
	}
	wg.Wait()
	close(resultChan)

	//считаем сумму из канала и выводим ее
	var result int
	for square := range resultChan {
		result += square
	}
	fmt.Printf("Sum of squares = %v", result)
}
