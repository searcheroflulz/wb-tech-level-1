package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]struct{})
	//добавляем элементы в мапу, где повторяющиеся просто не будут записаны
	for _, element := range sequence {
		m[element] = struct{}{}
	}
	var result []string

	for v := range m {
		result = append(result, v)
	}
	fmt.Println("Result:", result)
}
