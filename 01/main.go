package main

import "fmt"

var x int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

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

	p1 := person{
		"Miss",
		"Moneypenny",
	}
	p1.speak()

	sa1 := secretAgent{
		person:        person{
			"James",
			"Bond",
		},
		licenseToKill: true,
	}
	sa1.speak()
	sa1.person.speak()
}
