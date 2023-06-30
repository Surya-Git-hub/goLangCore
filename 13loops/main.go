package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "friday", "saturday"}

	fmt.Println("days")

	// for d:=0; d< len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:= range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		if rogueValue == 2 {
			goto lco
		}
		if rogueValue == 5 {
			break
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

back:
	for rogueValue < 10 {
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("jumping at lco ", rogueValue)
	if rogueValue < 10 {

		goto back
	}
}
