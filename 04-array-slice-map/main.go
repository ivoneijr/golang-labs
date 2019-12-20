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

func internalArraySlice() {
	slice1 := make([]int, 10, 20)
	slice2 := append(slice1, 1, 2, 3)

	fmt.Println(slice1, slice2)

	slice1[0] = 999
	fmt.Println(slice1, slice2)

}

func copySlice() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array = append(array1, 4,5,6)
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}

func map1() {
	// var approved map[int]string
	approved := make(map[int]string)

	approved[7263178637] = "Ivo"
	approved[4736874361] = "Nei"
	approved[6464672997] = "Jr"

	fmt.Println(approved)

	for cpf, nome := range approved {
		fmt.Printf("%s (CPF: %d\n", nome, cpf)
	}

	fmt.Println(approved[6464672997])
	delete(approved, 6464672997)
	fmt.Println(approved[6464672997])
}

func map2() {
	payments := map[string]float64{
		"José":  23444.12,
		"Maria": 55444.12,
		"Ivo":   123423444.12,
	}

	payments["Someone add"] = 1332.10
	delete(payments, "notExists")

	for name, payment := range payments {
		fmt.Println(name, payment)
	}
}

func map3() {
	byLetter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 14444.21,
			"Guga Pereira":   33343.34,
		},
		"J": {
			"José João": 2344.43,
		},
		"P": {
			"Pedro Paulo": 1200.00,
		},
	}

	delete(byLetter, "P")

	for letter, funcs := range byLetter {
		fmt.Println(letter, funcs)
	}
}

func main() {
	// array()
	// forRange()
	// slice()
	// sliceMake()
	// internalArraySlice()
	// copySlice()
	// map1()
	// map2()
	map3()
}
