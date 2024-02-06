package main

import (
	"fmt"
	"math"
)

type variable struct {
	x float64
	y float64
}

func add(param variable, z float64) float64 {
	return (param.x + param.y) * z
}
func subs(param *variable) float64 {
	return param.x - param.y

}
func square(param *variable) float64 {
	return math.Sqrt(param.x*param.x + param.y*param.y)

}

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("Input a variable")
	fmt.Scanf("%f", &a)
	fmt.Println("Input b variable")
	fmt.Scanf("%f", &b)
	fmt.Println("Input c variable")
	fmt.Scanf("%f", &c)
	v := variable{a, b}
	fmt.Println("Output a + b:")
	fmt.Println(add(v, c))
	fmt.Println("Output a - b:")
	fmt.Println(subs(&v))
	fmt.Println("Output square(a^2 + b^2):")
	fmt.Printf("result is %.2f\n", square(&v))

}
