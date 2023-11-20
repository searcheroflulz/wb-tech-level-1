package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	result := make(map[string]bool)

	for _, element := range sequence {
		result[element] = true
	}

	fmt.Println("Result:", result)
}
