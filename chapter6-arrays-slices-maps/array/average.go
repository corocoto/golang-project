package main

import "fmt"

func main() {
	values := [5]float64{
		25,
		65,
		13,
		45,
		21,
	}

	var total float64 = 0
	const valuesLength = len(values)
	for _, value := range values { /*
			It's like a `forEach` construction in JavaScript. First value contains item index, second - item, properly.
			If we don't use iterator (first value) - Golang will return the error on the compilation time.
			You need to change first element name in loop to `_` keyword for preventing this
		*/
		total += value
	}
	fmt.Println("Average is", total/float64(valuesLength))
}
