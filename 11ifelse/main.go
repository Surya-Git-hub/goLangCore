package main

import "fmt"

func main() {
	fmt.Println("if else in go lang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "bot"
	} else {
		result = "unique user"
	}

	fmt.Printf("result: %v\n", result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}

	// if err != nil {

	// }
}
