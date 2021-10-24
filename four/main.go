package main

import (
	"fmt"
	utility "four/util"
)

func main() {

	scores := [8]int{100, 80, 75, 92, 72, 93, 88, 67}
	result := utility.Avarage(scores)
	fmt.Println(result)

	resultSlice := utility.AddHightScore(scores)
	fmt.Println(resultSlice)
}
