package main

import "fmt"

func main() {
	// indirect
	var dataChicken = []map[string]string {
		map[string]string{"name" : "red", "gender" : "male"},
		map[string]string{"name" : "blue", "gender" : "female"},
		map[string]string{"name" : "yellow", "gender" : "male"},
	}

	for _, chicken := range dataChicken {
		fmt.Printf("Data gender %s have name %s \n", chicken["gender"], chicken["name"] )
	}
	// or
	var dataCostumer = []map[string]string {
		{"name" : "Tsukasa", "address" :"kyoto"},
		{"name" : "Yukinon", "address" :"Shibuya"},
		{"name" : "Isla", "address" :"Tokyo"},
	}

	for _, data := range dataCostumer {
		fmt.Printf("Costumer name %s came from %s \n", data["name"], data["address"])
	}
}
