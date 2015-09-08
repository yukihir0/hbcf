package hbcf

type SimilarityStrategy struct {
}

func NewSimilarityStrategy() SimilarityStrategy {
	return SimilarityStrategy{}
}

func (s SimilarityStrategy) CalculateScore(user User, item Item) float64 {
	return user.Similarity
}
