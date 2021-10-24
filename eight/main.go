package main

import (
	"eight/util"
	"fmt"
)

func main() {
	rectangle := util.Rectangle{20}
	large := util.CalculateLarge(rectangle)
	fmt.Println(large)

	sequare := util.Sequare{12, 45}
	largeSq := util.CalculateLarge(sequare)
	fmt.Println(largeSq)

}
