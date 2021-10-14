package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(fib(3)) // 2
	fmt.Println(fib(5)) // 5
	fmt.Println(fib(7)) // 13
}
