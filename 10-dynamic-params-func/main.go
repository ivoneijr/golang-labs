package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func printApproved(approved ...string) {
	fmt.Println("Approved list ")

	for i, appr := range approved {
		fmt.Printf("%d) %s\n", i+1, appr)
	}

}

func main() {

	list := []string{"Maria", "Pedro", "Ivonei", "Valentina", "Arianne"}

	printApproved(list...)

	// fmt.Println("Média %.2f", media(7.7, 8.1, 10.0))
	// fmt.Println("Média %.2f", media(7.0, 3.0))
}
