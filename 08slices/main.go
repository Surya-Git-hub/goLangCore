package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices in go")

	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is : %T \n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 3423
	highscore[2] = 344
	highscore[3] = 43
	// highscore[4] = 564

	highscore = append(highscore, 555, 6454, 3455)

	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))
	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))
}
