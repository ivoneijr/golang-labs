package main

import "fmt"

func justPrint() {
	fmt.Println("F1 justPrint")
}

func withArgs(first string, second string) {
	fmt.Printf("F2: %s %s \n", first, second)
}

func returnString() string {
	return "F3: returnString"
}

func paramAlias(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func doubleReturn() (string, string) {
	return "F5: first", "F5: second"
}

func main() {
	justPrint()
	withArgs("first", "second")
	returnString()
	fmt.Println(paramAlias("first", "second"))
	doubleReturn()
}
