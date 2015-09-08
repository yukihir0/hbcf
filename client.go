package hbcf

import "runtime"

var empty struct{}

type Client struct {
	totalPages         int
	maxParallelRequest int
	api                api
}

func NewClient() Client {
	return Client{
		totalPages:         1,
		maxParallelRequest: runtime.GOMAXPROCS(runtime.NumCPU()),
		api:                NewHttpAPI(),
	}
}

func (c *Client) SetTotalPages(total int) {
	if total > 0 {
		c.totalPages = total
	}
}

func (c *Client) SetMaxParallelRequest(max int) {
	if max > 0 {
		c.maxParallelRequest = max
	}
}

func (c *Client) SetAPI(api api) {
	c.api = api
}

func (c Client) RequestUserItems(users Users) []UserItem {
	uiChan := make(chan []UserItem)
	limitChan := make(chan struct{}, c.maxParallelRequest)

	go func() {
		for _, user := range users {
			for page := 0; page < c.totalPages; page++ {
				select {
				case limitChan <- empty:
					go func(user User, page int) {
						uiChan <- c.api.RequestUserItems(user, page)
						<-limitChan
					}(user, page)
				}
			}
		}
	}()

	userItems := []UserItem{}
	for i := 0; i < len(users)*c.totalPages; i++ {
		userItems = append(userItems, <-uiChan...)
	}

	return userItems
}
