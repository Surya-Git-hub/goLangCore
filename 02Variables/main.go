package main

import "fmt"

const Logintoken string = "kjsdfkldjf"


func main() {
	var username string = "Surya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLooggedIn bool = true
	fmt.Println(isLooggedIn)
	fmt.Printf("Variable is of type: %T \n", isLooggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float64 = 255.45555550500505000505505
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var website = "github/surya-git-hub"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	numberOfUser := 300000
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(Logintoken);
	fmt.Printf("Variable is of type: %T \n", Logintoken)

}
