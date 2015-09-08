package hbcf

import "github.com/yukihir0/hbn"

type Users []User

func NewUsers(neighbors hbn.Neighbors) Users {
	users := Users{}
	for _, neighbor := range neighbors {
		users = append(users, NewUser(neighbor))
	}
	return users
}
