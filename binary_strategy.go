package hbcf

type BinaryStrategy struct {
}

func NewBinaryStrategy() BinaryStrategy {
	return BinaryStrategy{}
}

func (s BinaryStrategy) CalculateScore(user User, item Item) float64 {
	return 1.0
}
