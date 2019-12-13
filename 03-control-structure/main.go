package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printResult(value float64) {
	if value >= 7 {
		fmt.Println("Approved")
	} else {
		fmt.Println("Reproved")
	}
	fmt.Println("lol")
}

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func testRandom() {
	if i := random(); i > 5 {
		fmt.Println("you win")
	} else {
		fmt.Println("you lose")
	}
}

func loops() {
	i := 1

	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println("")

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	for {
		fmt.Println("print and sleep")
		time.Sleep(time.Second)
	}
}

func switching(number interface{}) string {
	switch number.(type) {
	// do something about type
	}

	switch number {

	case 10:
		fallthrough

	case 9:
		return "TOP"

	case 8, 7:
		return "humm, ok"

	case 6, 5, 4, 3:
		return "bad guy"

	default:
		return "OMG, what?"
	}
}

func main() {
	// printResult(7.3)
	// printResult(6.3)
	// printResult("asd") cannot use "asd" (type string) as type float64 in argument to printResult

	// testRandom()

	// loops()

	fmt.Println(switching(10))
	fmt.Println(switching(9))
	fmt.Println(switching(7))
	fmt.Println(switching(4))
	fmt.Println(switching(0))
}
