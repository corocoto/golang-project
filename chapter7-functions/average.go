package main

import "fmt"

func average(values []float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}
	return total / float64(len(values))
}

func main() {
	valuesList := []float64{
		45.45,
		13.45,
		67.33,
		16.789,
		324.43,
		0.545,
	}
	fmt.Println(average(valuesList))
}
