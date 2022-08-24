package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

func ExecuteVisitorExample() {
	rectangle := Rectangle{10.0, 20.0}
	square := Square{5}
	circle := Circle{10}

	shapeAreaCalculator := ShapeAreaCalculator{}
	rectangle.accept(&shapeAreaCalculator)
	square.accept(&shapeAreaCalculator)
	circle.accept(&shapeAreaCalculator)
}

// Visitor

type Visitor interface {
	visitForRectangle(*Rectangle)
	visitForSquare(*Square)
	visitForCircle(*Circle)
}

type ShapeAreaCalculator struct {
	area float64
}

func (f *ShapeAreaCalculator) visitForRectangle(r *Rectangle) {
	f.area = r.acSide * r.bdSide
	fmt.Printf("Area of %s is %v\n", r.getType(), f.area)
}

func (f *ShapeAreaCalculator) visitForSquare(s *Square) {
	f.area = s.side * s.side
	fmt.Printf("Area of %s is %v\n", s.getType(), f.area)
}

func (f *ShapeAreaCalculator) visitForCircle(c *Circle) {
	f.area = 3.14 * (c.radius * c.radius)
	fmt.Printf("Area of %s is %v\n", c.getType(), f.area)
}

// Shapes

type Shape interface {
	getType() string
	accept(Visitor)
}

type Rectangle struct {
	acSide float64
	bdSide float64
}

func (r *Rectangle) accept(v Visitor) {
	v.visitForRectangle(r)
}

func (r *Rectangle) getType() string {
	return "rectangle"
}

type Square struct {
	side float64
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "square"
}

type Circle struct {
	radius float64
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "circle"
}
