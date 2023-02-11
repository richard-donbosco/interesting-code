package main

import "fmt"

type gender string

const (
	male   gender = "male"
	female gender = "female"
	other  gender = "other"
)

type person struct {
	name   string
	age    int
	gender gender
}

func main() {
	p1 := person{name: "Richard", age: 27, gender: male}
	p1.print()

	var p2 *person
	p2 = new_person("Michael", 27, male)
	p2.print()

}

func new_person(name string, age int, gender gender) *person {
	p := person{name: name, age: age, gender: gender}
	return &p
}

func (p person) print() {
	fmt.Printf("Name:%s \n", p.name)
	fmt.Printf("Age:%d\n", p.age)
	fmt.Printf("Gender:%v\n", p.gender)
	fmt.Println("------------")

}
