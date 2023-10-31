package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(totalArea(&Circle{0, 0, 5}))

	multiShape := MultiShape{
		shapes: []Shape{
			Circle{0, 0, 5},
			Rectangle{0, 0, 10, 10},
		},
	}

	fmt.Println(totalArea(multiShape.shapes...))
	fmt.Println(multiShape.area())
}

type Shape interface {
	area() float64
}

// We can use interfaces as arguments
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type MultiShape struct {
	//We can use interfaces as fields
	shapes []Shape
}

// Area With this, we does need totalArea function
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// Circle There is no explicit declaration of intent
type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle There is no explicit declaration of intent
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
