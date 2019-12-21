package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

// * is a pointer, should pass a mem reference
func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")
	p.name = parts[0]
	p.lastName = parts[1]
}

func main() {
	ivo := person{"Ivo", "Junior"}
	fmt.Println(ivo.getFullName())

	ivo.setFullName("Ivonei Jr")
	fmt.Println(ivo.getFullName())
}
