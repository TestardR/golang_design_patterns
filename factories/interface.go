package factories

import "fmt"

type Person interface {
	sayHello()
}

type person struct {
	name string
	age  int
}

type tiredPerson struct {
	name string
	age  int
}

func (p *person) sayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old", p.name, p.age)
}

func (p *tiredPerson) sayHello() {
	fmt.Printf("Sorry, I'm too tired.")
}

func NewPerson(name string, age int) Person {
	if age > 100 {
		return &tiredPerson{name, age}
	}
	return &person{name, age}
}

func main() {
	p := NewPerson("James", 34)
	p.sayHello()
}
