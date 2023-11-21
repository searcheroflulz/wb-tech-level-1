package main

import "fmt"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, key int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == key {
			return mid
		} else if arr[mid] < key {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

func main() {
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 1
	index := binarySearch(sortedArr, key)

	fmt.Printf("%v is the %v element", key, index)
}
