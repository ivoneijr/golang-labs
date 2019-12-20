package main

import "fmt"

func getApproved(value int) int {
	defer fmt.Println("END")

	if value > 5000 {
		fmt.Println("High value")

		return value
	}

	fmt.Println("Low value")

	return value

}

func main() {
	fmt.Println(getApproved(6000))
	fmt.Println(getApproved(3000))
}
