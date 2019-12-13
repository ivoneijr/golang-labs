package main

import (
	"fmt"
	"math"
	"time"
)

func sum(a int, b int) int {
	return a + b
}

func print(value int) {
	fmt.Println(value)
}

func operators() {
	a := 3
	b := 2

	fmt.Println("a = 3, b = 4, then:")
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a*b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a|b|=", a%b)

	// bitwise
	fmt.Println("a&b=", a&b)
	fmt.Println("a|b=", a|b)
	fmt.Println("a^b=", a^b)

	c := 3.0
	d := 2.0
	fmt.Println("c = 3.0, d = 2.0, then:")
	fmt.Println("max", math.Max(c, d))
	fmt.Println("min", math.Min(c, d))
	fmt.Println("pow", math.Pow(c, d))
}

func assignment() {
	var b byte = 3

	fmt.Println(b)

	c := 3 //inference
	c += 3 // c = c + 3
	c -= 3 // c = c - 3
	c /= 2 // c = c / 3
	c *= 2 // c = c * 3
	c %= 2 // c = c % 3

	d, e := 1, 2
	fmt.Println("d, e := 1,2", d, e)

	d, e = e, d
	fmt.Println("d, e = e, d", d, e)
}

func relOperators() {
	fmt.Println("Strings:", "Banana" == "Banana")
	fmt.Println("3 != 2", 3 != 2)
	fmt.Println("3 < 2", 3 < 2)
	fmt.Println("3 > 2", 3 > 2)
	fmt.Println("3 <= 2", 3 <= 2)
	fmt.Println("3 >= 2", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("d1 := time.Unix(0,0)")
	fmt.Println("d2 := time.Unix(0,0)")
	fmt.Println("d1 == d2", d1 == d2)
	fmt.Println("d1 == d2", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	fmt.Println("p1 := Pessoa('João')")
	fmt.Println("p1 := Pessoa('João')")

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println("p1 == p2", p1 == p2)

}

func logicOperators(job1, job2 bool) (bool, bool, bool) {
	buyIphone := job1 && job2
	buyTv := job1 != job2
	buyIcecream := job1 || job2

	return buyIphone, buyTv, buyIcecream
}

func unaryOperators() {
	x := 1
	y := 2
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	x++
	fmt.Println("x++", x)

	x--
	fmt.Println("x--", x)
}

func notTernaryTest(value float64) string {
	if value > 6 {
		return "approved"
	}

	return "try again :)"
}

func pointers() {
	ivar := 1

	var pointer *int = nil // create nil pointer (memory address with empty value)

	pointer = &ivar // get p memory address
	*pointer++      // remover reference (getting value)
	ivar++

	// pointer++ not permitted ioerations over pointers

	fmt.Println(pointer, *pointer, ivar, &ivar)
}

func main() {
	result := sum(1, 5)
	print(result)

	operators()
	assignment()
	relOperators()

	buyIphone, buyTv, buyIcecream := logicOperators(true, true)
	fmt.Printf("buyIphone: %t, buyTv: %t, buyIcecream: %t, healty: %t", buyIphone, buyTv, buyIcecream, !buyIcecream)

	unaryOperators()

	fmt.Println(notTernaryTest(3))
	fmt.Println(notTernaryTest(9))

	pointers()
}
