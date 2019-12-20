package main

import "fmt"

func change(x, y int) (second int, first int) {
	second = y
	first = x

	return // clear return
}

func main() {
	x, y := change(0, 1)

	fmt.Println(x, y)
}
