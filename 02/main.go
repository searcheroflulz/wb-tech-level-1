package main

import (
	"fmt"
	"math"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает
значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func getSquare(wg *sync.WaitGroup, num int) {
	defer wg.Done()
	fmt.Printf("square of %v = %v\n", num, math.Pow(float64(num), 2))

}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range arr {
		wg.Add(1)
		go getSquare(&wg, num)
	}
	wg.Wait()
}
