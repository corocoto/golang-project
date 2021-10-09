package main

import "fmt"

func main() {
	slice1 := []float64{1, 2, 3}
	slice2 := make([]float64, 2)
	copy(slice2, slice1)
	fmt.Println("slice1 =", slice1) // [1 2 3]
	fmt.Println("slice2 =", slice2) // [1 2] because slice2 only can contain 2 elements
}
