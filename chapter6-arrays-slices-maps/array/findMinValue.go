package main

import (
	"fmt"
	"sort"
)

func main() {
	numbersArray := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	sort.Ints(numbersArray)
	fmt.Println("Minimal value from array is", numbersArray[0])
}
