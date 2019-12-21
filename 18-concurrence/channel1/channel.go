package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // write data to channel
	<-ch    // read data from channel

	ch <- 4
	fmt.Println(<-ch)
}
