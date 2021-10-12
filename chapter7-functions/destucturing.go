package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func main() {
	fmt.Println("First result of sum:", add(1, 2, 3, 4, 5))
	arr := []int{1, 3, 5, 6, 5467, 57}
	fmt.Println("Second result of sum:", add(arr...))
}
