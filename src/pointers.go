package main

import (
	"fmt"
)

func main() {
	var p *int
	if p != nil {
		fmt.Println("Value of p: ", *p)
	} else {
		fmt.Println("P pointer is nil")
	}

	// lets assign the pointer
	var v int = 43
	p = &v

	// test again
	if p != nil {
		fmt.Println("Value of p: ", *p)
	} else {
		fmt.Println("P pointer is nil")
	}

	// now with implicit typing
	var myFloat float64 = 43.25
	p1 := &myFloat
	fmt.Println("Float value is : ", *p1)

	// modify pointer value
	*p1 = *p1 * 2
	fmt.Println("New Float value is : ", *p1)
	fmt.Println("Original myFloat value is : ", myFloat)

}
