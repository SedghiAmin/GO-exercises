package main

import(
	"fmt"
)

type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
    Base, Height, Side1, Side2 float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

func (t Triangle) Area() float64 {
    return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
    return t.Base + t.Side1 + t.Side2
}

func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	
	shapes := []Shape{
		Rectangle{Width: 5 , Height: 8},
		Circle{Radius: 10},
		Triangle{Base: 2, Height: 5, Side1: 3, Side2: 4},
	}
	
	for _, shape:= range shapes {
		/* fmt.Println(
			shape.Area(),
			shape.Perimeter(),
		) */
		PrintShapeInfo(shape)
	}
}