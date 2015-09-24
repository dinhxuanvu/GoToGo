package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	one := person{"Alex", 22}
	fmt.Println(one)
	fmt.Println(one.name)
	fmt.Println(one.age)
	one.name = "Anderson"
	fmt.Println(one.name)
	two := &one
	fmt.Println(two)
	fmt.Println(two.age)

	fmt.Println(person{"Bob", 21})
	fmt.Println(person{name: "Davis", age: 25})
}
