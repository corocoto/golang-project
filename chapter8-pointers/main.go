package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0 //change not only argument, but also original value that send as parameter
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // 0
}
