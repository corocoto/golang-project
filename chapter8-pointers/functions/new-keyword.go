package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	value := new(int)
	one(value)
	fmt.Println(*value) // 1
}
