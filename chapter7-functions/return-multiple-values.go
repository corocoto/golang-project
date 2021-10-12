package main

import "fmt"

func getTwoInt() (int, int) {
	return 5, 6
}

func main() {
	a, b := getTwoInt()
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
