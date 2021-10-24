package main

import (
	"first/service"
	"fmt"
)

func main() {
	name := "Miyuko"
	fmt.Println("Hello, Watasi Tsukasa Chan Desu!, and my wife name is, ", name)

	// get for data.go
	info := GetData()
	fmt.Println(info)

	// get for package entity
	resultAdd := service.Add(10, 15)
	fmt.Println(resultAdd)

	resultAdd2 := service.Add(15, 15)
	fmt.Println(resultAdd2)

	resultMultiply := service.Multiply(10, 10)
	fmt.Println(resultMultiply)

}
