package main

import "fmt"

func main() {
	exampleMap := make(map[string]string)
	exampleMap["greeting"] = "Hello"
	exampleMap["goodbye"] = "Bye"
	delete(exampleMap, "goodbye")
	fmt.Println(exampleMap)
}
