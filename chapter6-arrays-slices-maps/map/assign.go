package main

import "fmt"

func main() {
	chemicalElements := map[string]map[string]string{ //short creating process of map
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	/*
	 `if` block return 2 values.
	 first - request result, second - answer to the question: "is request be successful?"
	*/
	if name, ok := chemicalElements["Se"]; ok {
		fmt.Println(name, ok)
	}

	if name, ok := chemicalElements["O"]; ok {
		fmt.Println(name["name"], name["state"], ok) // Oxygen gas true
	}
}
