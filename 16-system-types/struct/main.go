package main

import "fmt"

type product struct {
	name     string
	value    float64
	discount float64
}

// receiver
func (p product) promoValue() float64 {
	return p.value * (1 - p.discount)

}

func main() {
	// long =(
	var pencil product

	pencil = product{
		name:     "pencil",
		value:    1.79,
		discount: 0.05,
	}
	fmt.Println(pencil, pencil.promoValue())

	// smart
	tv := product{"Lg Smart tv", 1989.90, 0.10}
	fmt.Println(tv, tv.promoValue())
}
