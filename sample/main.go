package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/yukihir0/hbcf"
	"github.com/yukihir0/hbn"
)

func main() {
	userItems := ImportUserItemsNetwork("yukihir0")
	//userItems := ImportUserItemsFile("csv/user_items.csv")
	//ExportUserItemsFile(userItems, "csv/user_items.csv")

	strategy := hbcf.NewSimilarityStrategy()
	//strategy := hbcf.NewBinaryStrategy()
	//strategy := hbcf.NewBookmarkCountStrategy()
	//strategy := hbcf.NewTorusStrategy()
	items := hbcf.Recommend(strategy, userItems)

	for _, item := range items.Top(20) {
		fmt.Printf("- %s : %.3f\n", item.Title, item.Score)
	}
}

func ImportUserItemsNetwork(user string) []hbcf.UserItem {
	client := hbn.NewClient()
	client.SetTotalPages(2)
	client.SetMaxParallelRequest(10)

	bookmarks := client.RequestBookmarks(user)
	//bookmarks := client.RequestRelatedBookmarks(user)
	//bookmarks := client.RequestFavoriteBookmarks(user)
	//bookmarks := client.RequestHotEntryBookmarks()
	//bookmarks := client.RequestSearchBookmarks("golang")

	neighbors := client.SearchNeighbors(bookmarks)
	//sort.Sort(sort.Reverse(neighbors))
	excluded := neighbors.Exclude([]string{user})
	users := hbcf.NewUsers(excluded.Top(20))

	c := hbcf.NewClient()
	c.SetTotalPages(2)
	c.SetMaxParallelRequest(10)
	userItems := c.RequestUserItems(users)

	return userItems
}

func ImportUserItemsFile(path string) []hbcf.UserItem {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	userItems := []hbcf.UserItem{}
	reader := csv.NewReader(file)
	for {
		columns, err := reader.Read()
		if err == io.EOF {
			break
		}

		similarity, _ := strconv.ParseFloat(columns[1], 64)
		user := hbcf.User{
			Name:       columns[0],
			Similarity: similarity,
		}

		count, _ := strconv.Atoi(columns[4])
		item := hbcf.Item{
			Title:         columns[2],
			URL:           columns[3],
			BookmarkCount: count,
		}

		userItems = append(userItems, hbcf.UserItem{
			User: user,
			Item: item,
		})
	}

	return userItems
}

func ExportUserItemsFile(userItems []hbcf.UserItem, path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for _, ui := range userItems {
		err := writer.Write([]string{
			ui.User.Name,
			strconv.FormatFloat(ui.User.Similarity, 'f', 4, 64),
			ui.Item.Title,
			ui.Item.URL,
			strconv.Itoa(ui.Item.BookmarkCount),
		})
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()
}
