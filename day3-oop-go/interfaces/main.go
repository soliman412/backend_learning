package main
import "fmt"

type Speaker interface {
	Speak()
}
type Dog struct {
	Name string
}
func (d Dog) Speak() {
	fmt.Println(d.Name , ":Woof")
}
type Cat struct {
	Name string
}
func (c Cat) Speak() {
	fmt.Println(c.Name , ":Meow")
}
type Robot struct {
	Model string
}
func (r Robot) Speak() {
	fmt.Println(r.Model , ":beeb boop")
}
func MakeItTalk (s Speaker) {
	s.Speak()
}
func main (){
	speaker := []Speaker {
		Dog{Name: "Rex"},
		Cat{Name: "Mimi"},
		Robot{Model: "R2"},
	}
	for _, s := range speaker {
		s.Speak()
	}
	MakeItTalk(Dog{Name: "Buddy"})
}