package main

import (
	"fmt"
	"time"
)

func two34(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go two34(2, c)

	a := <-c
	fmt.Println(a)

	b := <-c
	fmt.Println(a, b)

	fmt.Println(a, b, <-c)
}
