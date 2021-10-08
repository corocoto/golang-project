package main

import "fmt"

func main() {
	fmt.Println(len("Hello world")) //get string length
	fmt.Println("Hello world"[1]) //get symbol (on bytes format) placed on first index
	fmt.Println("Hello " + "world")
}
