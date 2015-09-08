package hbcf

import "github.com/yukihir0/hbapi"

type HttpAPI struct {
}

func NewHttpAPI() api {
	return HttpAPI{}
}

func (api HttpAPI) RequestUserItems(user User, page int) []UserItem {
	params := hbapi.NewFeedParams(user.Name)
	params.SetPage(page)
	feed, err := hbapi.GetFeed(params)
	if err != nil {
		return []UserItem{UserItem{
			User: user,
			Item: Item{},
		}}
	}

	userItems := []UserItem{}
	for _, item := range feed.Items {
		userItems = append(userItems, UserItem{
			User: user,
			Item: Item{
				Title:         item.Title,
				URL:           item.Link,
				BookmarkCount: item.BookmarkCount,
			}})
	}

	return userItems
}
