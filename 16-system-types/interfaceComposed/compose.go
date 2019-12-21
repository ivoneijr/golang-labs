package main

import "fmt"

type sport interface {
	activeTurbo()
}

type luxury interface {
	autoPark()
}

type luxurySport interface {
	sport
	luxury
}

type bmw6 struct{}

func (b bmw6) activeTurbo() {
	fmt.Println("Turbo On")
}

func (b bmw6) autoPark() {
	fmt.Println("auto parking")
}

func main() {
	var b luxurySport = bmw6{}
	b.activeTurbo()
	b.autoPark()
}
