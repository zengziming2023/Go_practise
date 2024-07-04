package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Bob2"})
	fmt.Println(person{age: 18})
	fmt.Println(*newPerson("newPeron"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(*sp)

	sp.age = 51
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		name:   "Dog",
		isGood: true,
	}
	fmt.Println(dog)
}
