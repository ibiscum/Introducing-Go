package main

import "fmt"

func main() {
	//Maps are also sometimes called associative arrays, hash tables, or dictionaries.
	var x map[string]int //map[KEY_TYPE]VALUE_TYPE.

	x = make(map[string]int)
	fmt.Println(len(x)) //0

	x["key"] = 10
	fmt.Println(len(x)) //1

	fmt.Println(x["key"]) //10

	fmt.Println(x) //map[key:10]

	delete(x, "key")

	fmt.Println(x) //map[]

	x["AnotherKey"] = 5
	if name, ok := x["AnotherKey"]; ok {
		fmt.Println(name, ok)
	}

	//Short way to write a map
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
	}
	fmt.Println(elements)

	//Using maps as objects
	elementsAsObject := map[string]map[string]string{
		"H": map[string]string{ //Is not necessary declare the type
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
			"state": "gas",
		},
	}

	if el, ok := elementsAsObject["C"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println(elementsAsObject)
}
