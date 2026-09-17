package main

import "fmt"

func maps() {
	ages := make(map[string]int)

	ages["cole"] = 3
	ages["black"] = 5
	ages["cole2"] = 2

	fmt.Println("cole, black, and cole2 are", ages["cole"], ",", ages["black"], ",", ages["cole2"], ",", "years old respectively")

	ages["cole"] = 10
	fmt.Println("seven years later cole becomes", ages["cole"])

	age, check := ages["cole"]
	if check {
		fmt.Println("cole is:", age)
	} else {
		fmt.Println(age, "is not found")
	}

	fmt.Println(age + 1)

	delete(ages, "cole")

	age, check = ages["cole"]
	if check {
		fmt.Println("cole is:", age)
	} else {
		fmt.Println(age, "is not found")
	}
}
