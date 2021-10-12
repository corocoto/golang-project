package main

import "fmt"

func main() {
	score := 0
	increment := func() int {
		score++
		return score
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
