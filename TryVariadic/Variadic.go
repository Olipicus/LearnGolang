package main

import "fmt"

func main() {

	fmt.Println(avg(20, 30, 28, 38))

	listNum := []float64{20, 30, 28, 38}

	/* You Can't do this
	fmt.Println(avg(listNum))
	*/

	//But You Can do this
	fmt.Println(avg(listNum...))
	fmt.Println(avg([]float64{20, 30, 28, 38}...))

}

func avg(listNum ...float64) float64 { //variadic function
	sum := 0.0

	for _, num := range listNum {
		sum += num
	}

	return sum / float64(len(listNum))
}
