package main

import "fmt"

func variableModification(age *int) int {
	*age = *age + 10
	return *age
}

func run() {
	age := 10
	fmt.Println(variableModification(&age))
	fmt.Println(age)

	fmt.Println(&age)
	age = 10
	fmt.Println(*&age)

	num := 7
	addr := &num
	fmt.Println(*addr)
}
