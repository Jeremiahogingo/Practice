package main

import (
	"fmt"
)

func main() {
	// Declare variables
	var str string
	var m, n int
	var mn float32

	// Assign values
	str = "Hello World!"
	m = 10
	n = 50
	mn = 2.45

	fmt.Println("Value of str = ", str)
	fmt.Println("Value of m = ", m)
	fmt.Println("Value of n = ", n)
	fmt.Println("Value of mn = ", mn)

	// Declare and assign values to variables
	var city string = "London"
	var x int = 100

	fmt.Println("Value of city = ", city)
	fmt.Println("Value of x = ", x)

	// Declare variable with defining its type
	country := "Kenya"
	number := 45

	fmt.Println("Value of country = ", country)
	fmt.Println("Value of number = ", number)

	// Declare multiple variables
	var (
		name  string
		email string
		age   int
	)
	name = "Jerrylion"
	email = "lion@gmail.com"
	age = 21

	fmt.Println("My name is , ", name)
	fmt.Println("Email me using ", email)
	fmt.Printf("I,m %d years old", age)
}
