package main

import "fmt"

func main() {
	printFizzBuzz(15)
}

func printFizzBuzz(num int) {
	if num%15 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}
