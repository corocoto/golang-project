package main

import "fmt"

func main() {
	exampleMap := make(map[string]int)
	exampleMap["twenty"] = 20
	fmt.Println(exampleMap)
	fmt.Println(exampleMap["twenty"])
}
