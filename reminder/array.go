package main

import "fmt"

func main() {
	var name[4] string
	name[0] = "Tsukasa chan"
	name[1] = "kaguya sama"
	name[2] = "Shino Aki"
	name[3] = "Yukinoshita"
	//names[4]= "test"

	for i, v := range name {
		fmt.Println("index - ",i,"=",v)
	}

	var fruits = [4]string {"Pisang", "Anggur", "Melon", "Pear"}
	fmt.Println("Panjang data buah: \t", len(fruits))
	fmt.Println("Data buah: \t", fruits)

	var animals = [5] string {
		"kadal",
		"badak",
		"rusa",
		"kucing",
	}
	for id, animal := range animals {
		fmt.Printf("Animal index %d  = %s \n", id, animal)
	}

	var value = [...]int {19, 12, 27, 12, 23, 37, 35}
	fmt.Println("Value of variable: ", value)
	fmt.Println("Lenght of Value: ", len(value))

	var no = [2][3]int {
		{12,14},
		{16,18},
	}

	fmt.Println(no)
	fmt.Println(len(no))

	for i := 0; i < len(no); i++ {
		for j := 0; j < len(no); j++ {
			fmt.Printf("array[%d][%d] = %d\n", i,j, no[i][j])
		}
	}

	var anime = make([]string, 2)
	anime[0] = "Kaguya sama, Love is War"
	anime[1] = "Sword Art Online"

	for i := 0; i < len(anime); i++ {
		fmt.Printf("Anime index %d = %s \n", i+1, anime[i])
	}
}
