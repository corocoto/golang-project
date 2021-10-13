package main

import "fmt"

func greeting() {
	fmt.Println("Hi")
}

func myNameIs() {
	fmt.Println("My name is Artem")
}

func main() {
	defer myNameIs()
	greeting()

	/*
		Output will be next:
		`Hi`
		`My name is Artem`

		Because we're using `defer` keyword, and it deferred call of `myNameIs` function
	*/
}
