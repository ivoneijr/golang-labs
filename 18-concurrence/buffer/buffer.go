package main

import (
	"fmt"
)

func routine(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("exec 4")
	c <- 5
	c <- 6
}
func main() {
	c := make(chan int, 3)

	go routine(c)

	fmt.Println(<-c)
	fmt.Println("done")
}
