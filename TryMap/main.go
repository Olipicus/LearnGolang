package main

import "fmt"

func main() {
	//Don't do like this
	//var myMap map[string]string
	//Because this map will nil

	myMap := map[string]string{} //By Short Identifier
	myMap["en"] = "Hello"
	myMap["fr"] = "Bonjour"
	myMap["th"] = "สวัสดี"
	fmt.Printf("%T\t%v\t%p\n", myMap, myMap, myMap)

	myMap = make(map[string]string) // By Make Function
	myMap["en"] = "Hello"
	myMap["fr"] = "Bonjour"
	myMap["th"] = "สวัสดี"
	fmt.Printf("%T\t%v\t%p\n", myMap, myMap, myMap)

	myMap = map[string]string{
		"en": "Hello",
		"fr": "Bonjour",
		"th": "สวัสดี",
	}
	fmt.Printf("%T\t%v\t%p\n", myMap, myMap, myMap)

	myMap["en"] = "Hi" // Modify Value By Key
	fmt.Printf("%T\t%v\t%p\n", myMap, myMap, myMap)

	delete(myMap, "fr") // Delete Value By Key
	fmt.Printf("%T\t%v\t%p\n", myMap, myMap, myMap)

}
