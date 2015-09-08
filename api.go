package hbcf

type api interface {
	RequestUserItems(user User, page int) []UserItem
}
