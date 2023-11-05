package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42
	return &p
}

func structImpl() {

	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{age: 30, name: "Alice"})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 40})
	fmt.Println(newPerson("john"))

	s := Person{name: "Sean", age: 50}
	fmt.Println(s)

	s.age = 36
	fmt.Println(s)

	sp := &s
	fmt.Println(sp)

	sp.age = 60
	fmt.Println(sp)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)

}
