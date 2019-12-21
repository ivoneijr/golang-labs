package main

import "fmt"

type car struct {
	name     string
	velocity int
}

type ferrari struct {
	car
	turbo bool
}

func main() {
	f := ferrari{}
	f.name = "enzo"
	f.velocity = 0
	f.turbo = true

	fmt.Printf("Ferrari is %s and the turbo is %v\n", f.name, f.turbo)
	fmt.Println(f)
}
