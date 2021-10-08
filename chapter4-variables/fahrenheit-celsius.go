package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter degrees by Fahrenheit: ")
	var fahrenheitDegrees float32
	fmt.Scanf("%f", &fahrenheitDegrees)
	celsiusDegrees := (fahrenheitDegrees - 32) * 5 / 9
	fmt.Println("Degrees by Celsius:", celsiusDegrees)
}
