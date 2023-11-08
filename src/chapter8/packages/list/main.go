package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	//Can be created with:
	//x := list.New()

	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	//Accept any type
	x.PushFront("Start")
	x.PushBack("End")

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) //Accept conversion to any type like e.Value.(int)
	}
}
