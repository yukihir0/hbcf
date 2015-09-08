package hbcf

type Item struct {
	Title         string
	URL           string
	BookmarkCount int
	Score         float64
}

func (item *Item) Update(score float64) Item {
	item.Score += score
	return *item
}
