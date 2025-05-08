package main

import (
	"fmt"
	"math"

	"github.com/LexusEgorov/mymath"
)

const PI = 3.141592

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shapes := make([]Shape, 0)
	areas := make([]float64, 0)

	shapes = append(shapes, Circle{Radius: 5}, Rectangle{Width: 4, Height: 6})

	for _, shape := range shapes {
		area := shape.Area()
		fmt.Println(area)

		areas = append(areas, area)
	}

	fmt.Printf("Avg area: %.2f", mymath.Average(areas))
}
