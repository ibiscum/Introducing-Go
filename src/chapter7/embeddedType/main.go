package main

import "fmt"

func main() {
	android := new(Android)
	android.Name = "C3PO"
	android.Talk()

	spaceship := new(Spaceship)
	spaceship.Commander.Name = "Luke"
	spaceship.Commander.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Spaceship struct {
	//This works like composition
	Commander Person // has a person
}

type Android struct {
	//This is an embedded type
	//This works like inheritance, but it is a composition yet
	Person // is a person
}
