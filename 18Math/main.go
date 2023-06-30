package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("The sum is : ", myNumberOne+int(myNumberTwo))

	//random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5)+1)

	myRandomNum, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println(myRandomNum)
}
