package main

import (
	"fmt"
)

func main() {
	defer func() {
		str := recover() // handle the exception by `recover` function
		fmt.Println(str) // [text from that sends to `panic` function]
	}() //IIFe
	panic("Error") // throw exception
}
