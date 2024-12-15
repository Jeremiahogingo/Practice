package main

import "fmt"

func main() {
	var a, b int

	// Assign values
	a = 5
	b = 10

	//Arithmetic operation
	// Addition
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)

	//subtraction
	d := a - b
	fmt.Printf("%d - %d = %d \n", a, b, d)

	//Multiplication
	e := a * b
	fmt.Printf("%d * %d = %d \n", a, b, e)

	// Division
	f := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n", a, b, f)

}
