package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")
var fruitList [4]string
fruitList[0] = "apple"
fruitList[1] = "Tomato"
fruitList[3] = "Peach"

fmt.Println("Fruit list is ",fruitList)
fmt.Println("Fruit list is ",len(fruitList))

var vegList = [5]string{"potato","beans","mushroom"}
fmt.Println("vegy list is : ", len(vegList))

}
