package main

import "fmt"

func half(number int) (int, bool) {
	halfOfNumber := number / 2
	isNumberParamEvent := number%2 == 0
	return halfOfNumber, isNumberParamEvent
}

func main() {
	fmt.Println(half(5))
	fmt.Println(half(1))
	fmt.Println(half(2))
}
