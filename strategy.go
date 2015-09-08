package hbcf

type strategy interface {
	CalculateScore(user User, item Item) float64
}
