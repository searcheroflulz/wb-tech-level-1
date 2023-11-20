package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func writeData(input chan int, arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(input)

	for _, num := range arr {
		input <- num
	}
}

func multiply(dataChan chan int, output chan int, wg *sync.WaitGroup) {
	defer close(output)
	defer wg.Done()

	for num := range dataChan {
		output <- num * 2
	}
}

func showResult(dataChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range dataChan {
		fmt.Printf("Result is %v\n", num)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	inputChan := make(chan int)
	outputChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go writeData(inputChan, numbers, &wg)

	wg.Add(1)
	go multiply(inputChan, outputChan, &wg)

	wg.Add(1)
	go showResult(outputChan, &wg)

	wg.Wait()
}
