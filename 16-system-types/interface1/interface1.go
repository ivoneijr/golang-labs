package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

type product struct {
	name  string
	price float64
}

// interfaces are implamented implicity
func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var some printable = person{"Ivo", "Jr"}
	fmt.Println(some.toString())
	print(some)

	some = product{"Stella Artois", 8.99}
	fmt.Println(some.toString())
	print(some)

}
