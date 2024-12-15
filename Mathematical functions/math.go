package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2.4
	b := 1.6

	// Power
	c := math.Pow(a, 2)
	fmt.Printf("%.2f ^ %d = %.2f \n", a, 2, c)

	//sin
	c = math.Sin(a)
	fmt.Printf("sin(%.2f) = %.2f \n", a, c)

	//cos
	c = math.Cos(b)
	fmt.Printf("cos(%.2f) = %.2f\n", b, c)

	//square root of a * b
	c = math.Sqrt(a * b)
	fmt.Printf("Sqrt(%.2f * %.2f) = %.2f \n", a, b, c)

}
