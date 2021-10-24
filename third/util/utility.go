package utility

import "fmt"

func OddWord(word string) {
	for i, v := range word {
		if i%2 == 0 {
			fmt.Printf("Huruf genap ke %d adalah: %s \n", i, string(v))
		}
	}
}

func CheckVocal(word string) {
	for i, v := range word {
		temp := string(v)
		switch temp {
		case "a", "i", "u", "e", "o":
			fmt.Printf("Index %d, : %s \n", i, temp)
		}
	}
}

func VocalPrint(word string) {
	for i, v := range word {
		vString := string(v)
		if vString == "a" || vString == "i" || vString == "u" || vString == "e" || vString == "o" {
			fmt.Printf("Index %d, : %s \n", i, vString)
		}

	}
}
