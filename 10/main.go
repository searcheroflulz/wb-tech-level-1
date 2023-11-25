package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.04, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
*/

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		//округляем температуру вниз до ближайшего десятка
		roundedTemp := int(math.Floor(temp/10) * 10)

		//добавляем температуру в соответствующую группу
		groups[roundedTemp] = append(groups[roundedTemp], temp)
	}

	return groups

}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := groupTemperatures(temperatures)

	for group, temps := range result {
		fmt.Printf("Группа %d градусов: %v\n", group, temps)
	}
}
