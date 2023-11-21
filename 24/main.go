package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func (p Point) distanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(2, 15)

	distance := point1.distanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
