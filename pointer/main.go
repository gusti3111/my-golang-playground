package main

import "fmt"

func main() {
	var j int = 80
	var a *int = &j

	fmt.Printf(" value of a is %d  and address of pointer is %d", *a, a)

}
