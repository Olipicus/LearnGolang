package main

import "fmt"

func main() {
	myRune := 'A'
	fmt.Printf("Type\tRuneVal\tStrVal\tByte\n")
	fmt.Printf("%T\t%v\t%v\t%v\n", myRune, myRune, string(myRune), []byte(string(myRune)))

	myRune = 'ก' //Thai Lang
	fmt.Printf("Type\tRuneVal\tStrVal\tByte\n")
	fmt.Printf("%T\t%v\t%v\t%v\n", myRune, myRune, string(myRune), []byte(string(myRune)))

	myRune = 65 // Define By Integer
	fmt.Printf("Type\tRuneVal\tStrVal\tByte\n")
	fmt.Printf("%T\t%v\t%v\t%v\n", myRune, myRune, string(myRune), []byte(string(myRune)))

	myString := "ABC"
	fmt.Printf("%v\t%v\t%v\n", myString[0], myString[1], myString[2])
	fmt.Printf("%v\t%v\t%v\n", string(myString[0]), string(myString[1]), string(myString[2]))

}
