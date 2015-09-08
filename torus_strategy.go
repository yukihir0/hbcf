package hbcf

import "math"

type TorusStrategy struct {
}

func NewTorusStrategy() TorusStrategy {
	return TorusStrategy{}
}

func (s TorusStrategy) CalculateScore(user User, item Item) float64 {
	// TODO userのsimilarityの分布を調べる(ヒストグラム）
	return math.Exp(-math.Pow(user.Similarity-0.5, 2) / 0.025)
}
