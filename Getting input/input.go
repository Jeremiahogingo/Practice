package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\t\tCircle area calculation\t\t")
	fmt.Println("\t\t=+=+=+=+=+=+=+=+=+=+=+=\t\t")
	fmt.Println("\nEnter radius of the circle	")

	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("The area of the circle is %.2f", area)
}
