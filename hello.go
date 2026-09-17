package main

import "fmt"

func hello() {
	fmt.Println("Hello World")

	// var age int = 18
	// var name string = "Dorcas"

	age := 37
	name := "Buchi"

	fmt.Println("This is", name+". He is", age, "years old.")

	language := "Golang"
	version := 1.21

	fmt.Println("I am learning", language, version)
	fmt.Printf("I am learning %s %.2f\n", language, version)

	const Time = 3.61
	const ProjectName = "Blog"

	fmt.Printf("The %s project was completed in %.2f seconds\n", ProjectName, Time)

	var ints int = 57
	var floats float64 = 57.777
	var bools bool = true
	var strings string = "Hello"

	fmt.Println(ints, floats, bools, strings)
}
