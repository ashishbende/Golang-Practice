package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str1 := "Good Morning!"
	str2 := "Beautiful day"
	str3 := "See you later"
	// returns string length & error
	stringLength, err := fmt.Println(str1, str2, str3)

	if err == nil {
		fmt.Println("String length: ", stringLength)
	}

	// we can discard err variable by using _
	stringLength1, _ := fmt.Println(str1, str2)
	fmt.Println("Now string length is : ", stringLength1)

	//output my values
	myVar := 30
	myBoolean := true

	fmt.Printf("myVar value: %v\n", myVar)
	fmt.Printf("myBoolean value: %v\n", myBoolean)

	// convert number to float
	myFloat := float64(myVar)
	fmt.Printf("myFloat value: %.2f\n", myFloat)

	// we can print data type
	fmt.Printf("str1 is of type %T\n", str1)
	fmt.Printf("myVar is of type %T\n", myVar)
	fmt.Printf("myBoolean is of type %T\n", myBoolean)
	fmt.Printf("myFloat is of type %T\n", myFloat)

	// we can store printf result in a string as well

	result := fmt.Sprintf("myFloat is of type %T\n", myFloat)
	fmt.Println(result)

	// now lets get some value from console
	// using var keyword with explicit type declaration
	var input string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&input)
	fmt.Println("Nice to meet you, ", input, " !")

	// using bufio .. something similar like java reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What color do you like: ")
	// single quote is infered as a byte rather than a string
	// '\n' specifies read till we hit enter button
	color, _ := reader.ReadString('\n')
	fmt.Println("Oh, ", color, " is my favorite color too!")

	// now lets get some numbers
	fmt.Print("Can you give me some random number: ")
	// note I'm reusing color variable, so no ':=' implicit declaration
	color, _ = reader.ReadString('\n')
	// remove spaces
	color = strings.TrimSpace(color)
	// convert string to float
	num, err := strconv.ParseFloat(color, 64)
	if err == nil {
		fmt.Printf("Given number: %.2f", num)
	} else {
		fmt.Println(err)
	}

}
