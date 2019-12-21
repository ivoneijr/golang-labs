package main

import "fmt"

type course struct {
	name string
}

func main() {
	var something interface{}
	fmt.Println(something)

	something = 3
	fmt.Println(something)

	type dynamic interface{}
	var something2 dynamic = "VISH"
	fmt.Println(something2)

	something2 = true
	fmt.Println(something2)

	something2 = course{"GOLANG RULES"}
	fmt.Println(something2)
}
