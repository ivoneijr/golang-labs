package main

import (
	"fmt"
	"reflect"
)

func array() {
	var values [3]float64 // not dynamic size

	fmt.Println(values)

	values[0], values[1], values[2] = 1.3, 4.55, 0.1
	fmt.Println(values)

	total := 0.0

	for i := 0; i < len(values); i++ {
		total += values[i]
	}

	media := total / float64(len(values))
	fmt.Printf("average: %.2f\n", media)
}

// video 32
func forRange() {
	numbers := [...]int{1, 2, 3, 4, 5}

	for i, num := range numbers { // get index and value
		fmt.Printf("[%d] %d\n", i+1, num)
	}

	for _, val := range numbers { // ignore index variable
		fmt.Println(val)
	}
}

func slice() {
	array1 := [3]int{1, 2, 3} //array
	slice1 := []int{1, 2, 3}  //slice (dynamic size)

	fmt.Println(array1, slice1)
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	array2 := [5]int{1, 2, 3, 4, 5}
	slice2 := array2[1:3] // create a pointer (2,3) exclude index 3
	fmt.Println(array2, slice2)

	slice3 := array2[:2] // from index 0 to index 2
	fmt.Println(array2, slice3)
}

func sliceMake() {
	slice := make([]int, 10)
	slice[9] = 100
	fmt.Println(slice)

	slice = make([]int, 10, 20) // 20  is internal array size
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice)) // when you add more than capacity, slice size is duplicated
}

func main() {
	// array()
	// forRange()
	// slice()
	sliceMake()
}
