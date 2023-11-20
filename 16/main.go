package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func quickSort(unsorted []int, low int, high int) {
	if len(unsorted) <= 1 {
		return
	}
	if low < high {
		p := partition(unsorted, low, high)
		quickSort(unsorted, low, p)
		quickSort(unsorted, p+1, high)
	}
}

func partition(unsorted []int, low int, high int) int {
	mid := (low + high) / 2
	pivot := unsorted[mid]
	i := low
	j := high
	for {
		for unsorted[i] < pivot {
			i++
		}
		for unsorted[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		unsorted[i], unsorted[j] = unsorted[j], unsorted[i]
		i++
		j--
	}
}

func main() {
	unsortedArr := make([]int, 10)
	for i := range unsortedArr {
		unsortedArr[i] = rand.Intn(50)
	}
	fmt.Println(unsortedArr)
	quickSort(unsortedArr, 0, len(unsortedArr)-1)
	fmt.Println(unsortedArr)
}
