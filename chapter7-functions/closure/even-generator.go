package main

import "fmt"

func makeEvenGenerator() func() uint {
	value := uint(0)
	return func() (result uint) { // явно указываем, какое значение будем возвращать
		result = value
		value += 2
		return
	}
}

func main() {
	evenValueGenerate := makeEvenGenerator()
	fmt.Println(evenValueGenerate()) // 0
	fmt.Println(evenValueGenerate()) // 2
	fmt.Println(evenValueGenerate()) // 4
}
