package main

import (
	calculate "five/util"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	resultAdd := calculate.Add(scores)
	fmt.Println(resultAdd)

	result, err := calculate.Calculate(4, 8, ".")
	if err != nil {
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
