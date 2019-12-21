package main

import "fmt"

type item struct {
	id    int
	qt    int
	value float64
}

type order struct {
	id    int
	items []item
}

// Here total eg.
func (p order) total() float64 {
	total := 0.0

	for _, item := range p.items {
		total += item.value * float64(item.qt)
	}

	return total
}

func main() {
	firstOrder := order{
		id: 3,
		items: []item{
			item{
				id:    1,
				qt:    10,
				value: 5.0,
			},
		},
	}

	fmt.Printf("amount: %.1f", order.total(firstOrder))
}
