package main

import "fmt"

func main() {
	/*
	int
	float32
	bool
	string
	*/

	// Numeric Types
	a := 10 // short variable declaration(works only inside a function)
	fmt.Println(a)

	var x int = 20 // explicit type declaration
	fmt.Println(x)

	var b = 30
	fmt.Println(b)

	// String Type
	var message string = "Hello, Go!"
	fmt.Println(message)

	// Boolean Type
	var isGoFun bool = true
	fmt.Println(isGoFun)

	// a := 10           // int
	// a := 40.34        // float64
	// a := "Hello"      // string
	// a := true         // bool
	// a = false         // bool (reassigned)

	const p = 100
	fmt.Println(p)

}