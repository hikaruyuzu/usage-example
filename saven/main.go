package main

import (
	"fmt"
	"saven/utility"
)

func main() {
	Id := 10
	fmt.Println("Address old ID: ", &Id)
	fmt.Println("value old ID: ", Id)
	utility.ChangeID(&Id, 100)
	fmt.Println("Address new ID: ", &Id)

	fmt.Println("value new ID: ", Id)

	student := utility.Student{
		ID:   1,
		Name: "Miyuki Koishiro",
		GPC:  3.75,
	}
	fmt.Println(student.Name)

	// utility.Graduate(&student)

	// fmt.Println(student.Name)

	student.GraduateName()

	fmt.Println(student.Name)

	gamers := utility.Gamers{
		Name:  "Miyuko",
		Games: nil,
	}
	gamers.AddGames("Cat Mario")
	gamers.AddGames("Genshin Impact")
	gamers.AddGames("Sword Art Online")
	gamers.AddGames("Blue Protocol")
	gamers.AddGames("OSU")
	gamers.AddGames("Tower Fantasy")
	gamers.AddGames("Dragon Raja")

	for _, game := range gamers.Games {
		fmt.Println("Game name: ", game)
	}
	fmt.Println(gamers)

}
