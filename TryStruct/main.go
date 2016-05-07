package main

import "fmt"

type person struct {
	name     string
	lastname string
}

func (p person) speakName() {
	fmt.Println(p.name)
}

func (p person) speakLastname() {
	fmt.Println(p.lastname)
}

type father struct {
	person   // like inheritance
	children map[int]person
}

//like overwrite
func (f father) speakName() {
	fmt.Println("Father name : ", f.name)
}

func main() {
	myPerson := person{
		name:     "robert",
		lastname: "downy",
	}

	fatherJohn := father{
		person: person{
			name:     "father robert",
			lastname: "downy",
		},
		children: make(map[int]person),
	}

	fatherJohn.children[0] = person{
		name:     "richard",
		lastname: "gear",
	}
	fmt.Println(myPerson)
	myPerson.speakName()
	fmt.Println(fatherJohn)
	fatherJohn.speakName()
	fatherJohn.speakLastname()
}
