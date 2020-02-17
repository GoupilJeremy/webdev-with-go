package main

import "fmt"

var x int

func main() {
	x := 7
	fmt.Println(x)

	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

}
