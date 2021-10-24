package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	name := []string {"Tsukasa", "Kawaii"}
	printMessage("Hello", name)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(3, 11)
	fmt.Println(randomValue)
}

// function
func printMessage(message string, arr[]string)  {
	var setNameString = strings.Join(arr, " ")
	fmt.Println(message, setNameString)
}

//function return val
func randomWithRange(min, max int) int {
	value := rand.Int() % (max - min + 1) + min
	return value
}
