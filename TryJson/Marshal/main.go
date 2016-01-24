package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string
	Lastname string
}

type father struct {
	person
	ListSon map[string]string
}

func printToJSON(obj interface{}) {
	bs, _ := json.Marshal(obj)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}

func main() {
	myFather := father{
		person: person{
			Name:     "Robert",
			Lastname: "Joe",
		},
		ListSon: make(map[string]string),
	}
	myFather.ListSon["FirstSon"] = "Bookie"
	myFather.ListSon["SecondSon"] = "Robb"
	fmt.Println(myFather)

	printToJSON(myFather)

	myPerson := person{
		Name:     "Tong",
		Lastname: "Theerapat",
	}

	printToJSON(myPerson)

}
