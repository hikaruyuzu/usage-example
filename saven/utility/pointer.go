package utility

import "fmt"

func ChangeID(old *int, new int) {
	*old = new
	fmt.Println("Address Change ID: ", old)
	fmt.Println("value Change ID: ", *old)
}

func Graduate(student *Student) {
	student.Name = student.Name + " S.Com"
	fmt.Println("Change name: ", student.Name)
}

func (student *Student) GraduateName() {
	student.Name = student.Name + " S.Com"
	fmt.Println("Change: ", student.Name)
}

func (gamers *Gamers) AddGames(newGame string) {
	gamers.Games = append(gamers.Games, newGame)
}
