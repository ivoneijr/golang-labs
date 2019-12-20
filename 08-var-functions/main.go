package main

import "fmt"

var sum = func(a, b int) int {
	return a + b
}

func main() {
	result := sum(1, 1)

	fmt.Println(result)
}
