package main

import "fmt"

type sport interface {
	activeTurbo()
}

type ferrari struct {
	model    string
	turbo    bool
	velocity int
}

func (f *ferrari) activeTurbo() {
	f.turbo = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.activeTurbo()

	// var car2 sport = ferrari{"F40", false, 0}
	var car2 sport = &ferrari{"F40", false, 0}
	car2.activeTurbo()

	fmt.Println(car1, car2)
}
