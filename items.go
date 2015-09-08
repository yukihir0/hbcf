package hbcf

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items Items) Less(i, j int) bool {
	var ret bool
	if items[i].Score == items[j].Score {
		ret = items[i].Title < items[j].Title
	} else {
		ret = items[i].Score > items[j].Score
	}
	return ret
}

func (items Items) Top(max int) Items {
	return items[0:max]
}
