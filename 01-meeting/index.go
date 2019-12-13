package main

import (
	"fmt"
	"math"
)

func simple_print() {
	fmt.Println("Hello")
	fmt.Println("Hello")
}

func const_vars() {
	const PI float64 = 3.1415
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("Area ", area)

	const (
		a = 1
		b = 2
	)

	fmt.Println("a ", a, "b ", b)

	const c, d bool = true, false

	fmt.Println(c, d)

	e, f, g := 2, false, "EITA"

	fmt.Println(e, f, g)
}

func main() {
	simple_print()
	const_vars()
}
