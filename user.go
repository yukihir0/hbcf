package hbcf

import "github.com/yukihir0/hbn"

type User struct {
	Name       string
	Similarity float64
}

func NewUser(neighbor hbn.Neighbor) User {
	return User{
		Name:       neighbor.User,
		Similarity: neighbor.Similarity,
	}
}
