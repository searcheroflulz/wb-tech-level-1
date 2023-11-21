package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func deleteElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = deleteElement(slice, 3)
	fmt.Println(slice)
}
