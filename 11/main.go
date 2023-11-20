package main

import (
	"fmt"
	"math/rand"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func intersection(set1, set2 []int) []int {
	result := make([]int, 0)

	// Создаем мапу для проверки уникальности элементов
	seen := make(map[int]bool)

	// Добавляем элементы из первого множества в мапу
	for _, elem := range set1 {
		seen[elem] = true
	}

	// Проверяем наличие элементов из второго множества в мапе
	for _, elem := range set2 {
		if seen[elem] {
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	numbers1 := make([]int, 10)
	for i := range numbers1 {
		numbers1[i] = rand.Intn(30)
	}

	numbers2 := make([]int, 10)
	for i := range numbers2 {
		numbers2[i] = rand.Intn(30)
	}

	result := intersection(numbers1, numbers2)

	fmt.Println("Intersections are:", result)

}
