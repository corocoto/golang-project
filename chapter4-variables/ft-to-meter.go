package main

import "fmt"

func main() {
	fmt.Print("Enter ft value: ")
	var ftValue float32
	fmt.Scanf("%f", &ftValue)
	meterValue := ftValue / 3.281
	fmt.Println("Meters value:", meterValue)
}
