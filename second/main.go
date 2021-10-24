package main

import "fmt"

func allowedOrNot(age int) {
	if age > 10 {
		fmt.Println("Boleh bermain game")
	} else {
		fmt.Println("Belum boleh bermain")
	}
}

func gradee(score int) string {

	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "B"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	return grade
}

func checkMember(username string) {
	if username != "" {
		if username == "Pekora" {
			fmt.Printf("Welcome back %s \n", username)
		} else {
			fmt.Printf("%s Dare Deska \n", username)
		}
	} else {
		fmt.Println("Please enter your username")
	}
}

// include switch
func add(val int, val1 int) int {
	return val + val1
}

func sub(val int, val1 int) int {
	return val - val1
}

func mul(val int, val1 int) int {
	return val * val1
}

func calculator(operation string, value int, value1 int) int {
	var temp int

	switch operation {
	case "1":
		temp = add(value, value1)
	case "2":
		temp = sub(value, value1)
	case "3":
		temp = mul(value, value1)
	default:
		fmt.Println("Please enter correct operation or value")
	}

	return temp
}

func checkLength(word string) {
	l := len(word)

	switch {
	case l < 5:
		fmt.Println("username terlalu pendek")
	case l > 10:
		fmt.Println("username terlalu panjang")
	case l >= 5 && l <= 10:
		fmt.Println("Username sudah tepat")

	default:
		fmt.Println("Input kata tidak boleh kosong")

	}
}

func main() {

	age := 11
	allowedOrNot(age)

	grade := gradee(50)
	fmt.Println(grade)

	checkMember("Pekora")

	result := calculator("3", 12917291, 1261526)
	fmt.Println("Hasil adalah: ", result)

	checkLength("miyuko")
}
