package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		otherFactorial, _ := factorial(n - 1)

		return n * otherFactorial, nil
	}
}

func factorial2(n uint) uint {
	if n == 0 {
		return 1
	}

	return n * factorial2(n-1)
}

func main() {
	// factorial 1
	result, _ := factorial(5)
	fmt.Println(result)

	_, err := factorial(-4)

	if err != nil {
		fmt.Println(err)
	}

	// factorial 2
	fmt.Println(factorial2(5))
}
