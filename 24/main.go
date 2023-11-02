package main

import (
	"fmt"
	"math"
)

// Создаем структуру Point
type Point struct {
	x, y float64
}

// Конструктор создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p1 *Point) DistanceTo(p2 *Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}
func main() {
	A := NewPoint(5.0, 12.0)
	B := NewPoint(0.0, 0.0)

	distance := A.DistanceTo(B)
	fmt.Printf("Расстояние: %.3f\n", distance)
}
