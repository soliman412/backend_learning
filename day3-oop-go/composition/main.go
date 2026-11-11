package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, "is eating")
}

func (a *Animal) Info() {
	fmt.Println(a.Name, "- age", a.Age)
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) Bark() {
	fmt.Println(d.Name, "says: Woof!")
}

func main() {
	rex := Dog{
		Animal: Animal{Name: "Rex", Age: 3},
		Breed:  "Husky",
	}

	rex.Eat()
	rex.Info()
	rex.Bark()
	fmt.Println(rex.Breed)
}