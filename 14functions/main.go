package main

import "fmt"

func main() {
	fmt.Println("welcome to the functions in go lang")
	fmt.Println(adder(1,2));
	fmt.Println(proadder(1,2,3,45,5))

}

func adder(value int, value2 int) int {
	return value + value2
}

func proadder(values ...int) (int, string) {
	var result int
	for _, value := range values {
		result = result + value
	}
	return result, "success"
}
