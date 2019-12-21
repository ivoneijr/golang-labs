package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)

			c <- fmt.Sprintf("%s falando: %d", person, i)
		}
	}()

	return c
}

func group(first, second <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-first:
				c <- s
			case s := <-second:
				c <- s
			}
		}
	}()

	return c
}

func main() {
	c := group(speak("Ivo"), speak("Valentina"))

	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
