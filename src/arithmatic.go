package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	num1, num2, num3 := 10, 20, 30
	intSum := num1 + num2 + num3
	fmt.Println(num1, " + ", num2, " + ", num3, " = ", intSum)

	f1, f2, f3 := 10.2, 20.2, 30.8
	floatSum := f1 + f2 + f3
	fmt.Println(f1, " + ", f2, " + ", f3, " = ", floatSum)

	var b1, b2, b3, bigSum big.Float

	b1.SetFloat64(10.34)
	b2.SetFloat64(22.46)
	b3.SetFloat64(30.40)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	//fmt.Println("Big Sum value: ", bigSum)
	fmt.Printf("Big sum %.10g\n", &bigSum)

	radius := 12.5
	circleCircumference := radius * math.Pi
	fmt.Println("Circumference is : ", circleCircumference)

}
