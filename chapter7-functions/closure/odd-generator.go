package main

import "fmt"

func makeOddGenerator() func() uint {
	value := uint(1)
	return func() (result uint) {
		result = value
		value += 2
		return
	}
}

func main() {
	oddValueGenerator := makeOddGenerator()
	fmt.Println(oddValueGenerator())
	fmt.Println(oddValueGenerator())
	fmt.Println(oddValueGenerator())
}
