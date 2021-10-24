package main

import (
	"fmt"
	example "six/util"
)

func main() {

	user1 := example.User{}
	user1.ID = 1
	user1.FirstName = "Miyuki"
	user1.LastName = "Koishiro"
	user1.Email = "miyukikoishiro@netindo.com"
	user1.IsActive = true

	fmt.Println(user1)
	fmt.Println(user1.Email)

	user2 := example.User{
		ID:        2,
		FirstName: "Kaguya",
		LastName:  "Shinomiya",
		Email:     "shinomiyakaguya@gmail.com",
		IsActive:  true,
	}
	fmt.Println(user2)

	user3 := example.User{
		3,
		"Tsukasa",
		"Chan",
		"tsukasachan@gamial.com",
		false,
	}
	fmt.Println(user3)

	user4 := example.User{4, "Reina", "Yuzaki", "reina@jp.com", true}

	viewDisplay1 := example.ViewDisplay(user4)
	viewDisplay2 := example.ViewDisplay(user2)

	fmt.Println(viewDisplay1)
	fmt.Println(viewDisplay2)

	v, err := user4.Verivication("Reina Shinomiya", "reina@jp.com")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(v)
	}

	// Display user
	users := []example.User{user2, user4}

	group := example.Group{
		Name:        "Sagiri Wangy",
		Admin:       user4,
		Users:       users,
		IsAvailable: true,
	}

	group.DisplayGroup()

	dataUser4 := user4.DisplayUser()
	fmt.Println(dataUser4)

	dataUser2 := user2.DisplayUser()
	fmt.Println(dataUser2)
}
