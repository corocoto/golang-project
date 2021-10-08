package main

import "fmt"

func main() {
	fmt.Println("true && true =", true && true) // logical "AND"
	fmt.Println("true && false =", true && false)
	fmt.Println("true || true =", true || true) // logical "OR"
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true) // logical "NOT"
	fmt.Println("!false =", !false)
}
