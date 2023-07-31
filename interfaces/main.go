package main

import (
	"fmt"
	"math"
)

type equations2d interface {
	area() float64
	perimeter() float64
}
type equations3d interface {
	volume() float64
}
type circle struct {
	radius float64
}
type rect struct {
	width, heigth float64
}
type cube struct {
	edge float64
}
type cylinder struct {
	radius float64
}

// create method for area rectangle
func (r rect) area() float64 {
	return r.width * r.heigth

}

// create method for perim rectangle
func (r rect) perimeter() float64 {
	return 2*r.heigth + 2*r.width

}

// create method for area circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// create method for perim circle
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius

}

// create method for volume cube
func (c cube) volume() float64 {
	return c.edge * c.edge * c.edge

}

// create volume cylinder
func (c cylinder) volume() float64 {
	return math.Pi * c.radius * c.radius

}
func calculate2d(e2d equations2d) {
	fmt.Println(e2d)
	fmt.Println(e2d.area())
	fmt.Println(e2d.perimeter())

}
func calculate3d(e3d equations3d) {
	fmt.Println(e3d)
	fmt.Println(e3d.volume())
}
func main() {
	r := rect{width: 14, heigth: 51}
	circ := circle{radius: 5}
	cubical := cube{edge: 5}
	cyl := cylinder{radius: 49}

	calculate2d(r)
	calculate2d(circ)
	calculate3d(cubical)
	calculate3d(cyl)

}
