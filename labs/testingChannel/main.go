package main

import (
	"bytes"
	"fmt"
	"strings"
)

func makeFlat(matrix [][]string) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer close(ch)

		for _, row := range matrix {
			for _, column := range row {
				ch <- column
			}
		}
	}()

	return ch
}

func main() {
	matrix := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	flatted := makeFlat(matrix)

	var buffer bytes.Buffer

	for n := range flatted {
		buffer.WriteString(fmt.Sprintf("%s,", n))
	}

	response := strings.TrimSuffix(buffer.String(), ",")
	fmt.Println(response)
}
