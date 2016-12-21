package main

// factored import statements
import(
	"fmt"
	"math"
	"math/rand"
)

// functions with 2 arguments
func add(x int, y int) int{
	return x+y
}

// if parameters share the same type, put it only at the last
func sub(x,y int)int{
	return x-y
}

// multiple returns
func swap(x,y int)(int,int){
	return y,x
}

// the code starts here
func main(){
	fmt.Println("Hello Go! ")
	/* Exported name : Pi
	   In Go, any name with a CAPITAL letter, following the package name
	   is an exported name.
	*/
	fmt.Println("Random number is:",rand.Intn(10))
	fmt.Println(math.Pi)
	fmt.Println("Addition: ",add(2,3))
	fmt.Println("Substraction: " , sub(5,2))
	// multi value swaps. This reminds me of Python swap function
	a,b :=swap(5,2)
	fmt.Println("Swap: a= ",a," and b= ",b)
}

