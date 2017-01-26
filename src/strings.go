package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "Implicitly typed string"
	fmt.Printf("My String %v : %T\n", myString, myString)
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))

	string1 := "Hello"
	string2 := "hello"
	fmt.Println("string1 and string2 equal? ", (string1 == string2))
	fmt.Println("string1 and string2 Non-case sensitive equal? ", strings.EqualFold(string1, string2))

	fmt.Println("string1 Contains 'llo'?", strings.Contains(string1, "llo"))
}
