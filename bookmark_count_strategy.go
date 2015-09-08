package hbcf

type BookmarkCountStrategy struct {
}

func NewBookmarkCountStrategy() BookmarkCountStrategy {
	return BookmarkCountStrategy{}
}

func (s BookmarkCountStrategy) CalculateScore(user User, item Item) float64 {
	return float64(item.BookmarkCount)
}
