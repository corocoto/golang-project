package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		str := ""
		if i%3 == 0 {
			str = "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if len(str) == 0 {
			str = strconv.Itoa(i)
		}
		fmt.Println(str)
	}
}
