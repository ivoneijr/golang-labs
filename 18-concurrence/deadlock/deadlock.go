package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("after read")
}
func main() {
	c := make(chan int) // no buffer

	go routine(c)

	fmt.Println(<-c) // block
	fmt.Println("done")
}
