package main

import "fmt"

func swap(firstNumber *int, secondNumber *int) {
	*firstNumber, *secondNumber = *secondNumber, *firstNumber
}
func main() {
	x := 1
	y := 2
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	swap(&x, &y)
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
