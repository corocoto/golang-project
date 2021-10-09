package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}       // [1,2,3]
	slice2 := append(slice1, 4, 5) // [1,2,3,4,5]
	fmt.Println("slice2:", slice2)
}
