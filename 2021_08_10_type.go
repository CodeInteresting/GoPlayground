package main

import "fmt"

type B interface {
	ScaleBy(f float64)
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Sum() float64 {
	return p.X + p.Y
}


func main() {

	point := Point{.5, .6}
	fmt.Println(point)
	fmt.Println(point.Sum())
	point.ScaleBy(2.)
	fmt.Println(point)

	var x interface{}
	x = &Point{3.0, 0.8}
	_ = x.(B)

}