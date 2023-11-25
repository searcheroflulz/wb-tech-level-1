package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func deleteElement(slice []int, index int) []int {
	//если i ниже нуля или больше размера слайса, просто возвращаем слайс обратно
	if index < 0 || index >= len(slice) {
		return slice
	}
	//возвращаем новый слайс, в котором соединяем все элементы из начала прошлого до i + все от i до конца оригинального слайса
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = deleteElement(slice, 3)
	fmt.Println(slice)
}
