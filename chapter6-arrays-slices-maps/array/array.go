package main

import "fmt"

func main() {
	var numerableArray [5]int // creates array that contains 5 empty items
	numerableArray[4] = 100
	fmt.Println(numerableArray) // [0 0 0 0 100]
}
