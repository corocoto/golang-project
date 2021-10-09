package main

import "fmt"

func main() {
	var slice1 []float64           // creates empty slice
	fmt.Println("slice1:", slice1) // []

	slice2 := make([]float64, 5)   // creates slice which can contain 5 elements
	fmt.Println("slice2:", slice2) // [0 0 0 0 0]

	slice3 := make([]float64, 5, 10) // third argument is array length that slice link
	fmt.Println("slice3:", slice3)

	array := [5]float64{1, 2, 3, 4, 5}
	slice4 := array[2:4] //[3 4]
	fmt.Println("slice4:", slice4)
}
