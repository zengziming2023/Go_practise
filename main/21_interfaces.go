package main

import (
	"fmt"
	"math"
	"reflect"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect2 struct {
	width, height float64
}

type cicle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}

func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c cicle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c cicle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	t := reflect.TypeOf(g)
	fmt.Println(t.Name(), t.String())
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect2{width: 3, height: 4}
	c := cicle{radius: 5}

	interfaceType := reflect.TypeOf((*geometry)(nil)).Elem()
	rect2Type := reflect.TypeOf(r)
	iml := rect2Type.Implements(interfaceType)
	fmt.Println("rect2Type.Implements(interfaceType): ", iml)
	measure(r)
	measure(c)
}
