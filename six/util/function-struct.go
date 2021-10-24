package example

import (
	"errors"
	"fmt"
)

func ViewDisplay(user User) string {
	result := fmt.Sprintf("Name: %s %s Email: %s", user.FirstName, user.LastName, user.Email)
	return result
}

func (user User) Verivication(username, email string) (string, error) {
	fullName := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	if fullName == username && user.Email == email {
		result := fmt.Sprintf("Welcome Back %s \n", username)
		return result, nil
	} else {
		err := errors.New("Can't find username " + username)
		return "", err
	}
}

func (user User) DisplayUser() string {
	return fmt.Sprintf("Name: %s %s Email: %s", user.FirstName, user.LastName, user.Email)
}

func (group Group) DisplayGroup() {
	fmt.Printf("Name: %s \n", group.Name)

	MembersCount := len(group.Users)
	fmt.Printf("Members Count: %d \n", MembersCount)
	fmt.Printf("Users name: \n")

	for _, users := range group.Users {
		fmt.Printf("%s %s \n", users.FirstName, users.LastName)
	}

	fmt.Printf("Admin %s %s", group.Admin.FirstName, group.Admin.LastName)
}
