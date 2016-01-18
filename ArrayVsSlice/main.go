package main

import "fmt"

func main() {
	arrayA := [10]int{1, 2, 5, 1, 0}
	sliceA := []int{1, 2, 3, 4, 5}

	fmt.Println("ArrayType\tSliceType")
	fmt.Printf("%T\t\t%T\n", arrayA, sliceA)

	fmt.Println("------------- Explore Array -------------")
	fmt.Printf("value : %v\t size : %v\n", arrayA, len(arrayA))
	fmt.Println("-----------------------------------------")

	fmt.Println("------------- Explore Slice -------------")
	fmt.Printf("value : %v\t\t size : %v\n", sliceA, len(sliceA))
	fmt.Println("-----------------------------------------")

	fmt.Println(arrayA[1:3])

	sliceB := make([]int, 0, 5)
	fmt.Printf("%p\n", &sliceB)

	for _, val := range []int{1, 2, 3, 4, 5} {
		sliceB = append(sliceB, val)
	}
	fmt.Printf("%p\n", &sliceB)

	sliceC := sliceB
	fmt.Printf("%p\n", sliceC)

	sliceD := []int{}
	fmt.Println(sliceD == nil)

	//	sliceC := sliceB

}
