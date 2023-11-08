package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

// Len is the number of elements in the collection
func (ps ByName) Len() int {
	return len(ps)
}

// Less is the method that defines the sorting order
func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

// Swap swaps the elements
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Jill", 9},
		{"Jack", 10},
	}

	//The Sort function in sort package accepts any type that implements the sort.Interface
	//The sort.Interface requires three methods: Len, Less, and Swap
	sort.Sort(ByName(kids))

	fmt.Println(kids)
}
