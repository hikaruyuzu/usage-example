package main

import "fmt"

func main() {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["Januari"] = 40
	chicken["Februari"] = 90
	chicken["Maret"] = 80

	fmt.Println(chicken["Januari"])

	var name = make(map[string]string)
	name["Higehiro"] = "Sayu"
	name["Oregairu"] = "Yukinoshita"
	name["Tonikaku Kawai"] = "Tsukasa Chan"

	fmt.Println(name)
	fmt.Println(len(name))

	delete(name, "Higehiro")
	fmt.Println(name)
	fmt.Println(len(name))

	var fruits = map[string] int{}
	fruits["Januari"] = 50
	fruits["Maret"] = 90

	fmt.Println(fruits)

	var data = map[string]string {
		"name" : "Tsukasa Chan",
		"Country" : "Japan",
		"Age" : "19",
		"Status" : "Married",
	}

	for key, v := range data {
		fmt.Printf("Personal %s = %s \n", key, v)
	}

	// with pointer
	// food := *new(map[string]string) // return pointer data

	waifumu := map[string]string {
		"first" : "Tsukasa Chan",
		"Second" : "Kaguya sama",
		"Third" : "Yukinoshita",
	}

	var v, isExist = waifumu["first"]

	if isExist {
		fmt.Printf(v)
	} else {
		fmt.Printf("Key dont exist")
	}

}
