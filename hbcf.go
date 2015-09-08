package hbcf

import "sort"

func Recommend(s strategy, userItems []UserItem) Items {
	unique := map[string]Item{}
	for _, ui := range userItems {
		score := s.CalculateScore(ui.User, ui.Item) / float64(len(userItems))
		if i, ok := unique[ui.Item.URL]; ok {
			unique[ui.Item.URL] = i.Update(score)
		} else {
			unique[ui.Item.URL] = ui.Item
		}
	}

	items := Items{}
	for _, item := range unique {
		items = append(items, item)
	}
	sort.Sort(items)

	return items
}
