package main

import (
	"fmt"
	"sort"
)

/* abstraction to arrays
array is created behind the scene at runtime, resizable and can be sorted
much like Lists in python
*/
func main() {
	var mySlice = []string{"Red", "Green", "Blue"}
	fmt.Println(mySlice)
	// to add an item use append
	mySlice = append(mySlice, "Yellow")
	fmt.Println(mySlice)
	// to remove an item use append again
	// first item
	mySlice = append(mySlice[1:len(mySlice)])
	fmt.Println(mySlice)
	// last item
	mySlice = append(mySlice[:len(mySlice)-1])
	fmt.Println(mySlice)

	// build slice using make() function

	usingMake := make([]int, 5, 5)
	usingMake[0] = 0
	usingMake[1] = 5
	usingMake[2] = 3
	usingMake[3] = 8
	usingMake[4] = 1
	fmt.Println(usingMake, " Current capacity", cap(usingMake))
	// adding another item doubles the capacity
	usingMake = append(usingMake, 55)
	fmt.Println(usingMake, " Current capacity", cap(usingMake))

	sort.Ints(usingMake)
	fmt.Println("Sorted array", usingMake)

}
