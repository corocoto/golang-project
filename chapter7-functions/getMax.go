package main

import "fmt"

func getMaxNumber(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		if number > result {
			result = number
		}
	}
	return result
}

func main() {
	numsArray := []int{
		1,
		5,
		2,
		4,
		3,
		5,
		1,
	}
	maxNumberFromArray := getMaxNumber(numsArray...)
	fmt.Println(maxNumberFromArray)
}
