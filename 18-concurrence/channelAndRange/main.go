package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primes(n int, c chan int) {
	init := 2

	for i := 0; i < n; i++ {
		for prime := init; ; prime++ {
			if isPrime(prime) {
				c <- prime
				init = prime + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 30)

	go primes(cap(c), c)

	for i := range c {
		fmt.Printf("%d ", i)
	}

	fmt.Println("Done")
}
