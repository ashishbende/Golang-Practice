package main

import "fmt"

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	// get length
	fmt.Println("Number of items in array numbers: ", len(numbers))
}
