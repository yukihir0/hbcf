# hbcf

"hbcf" is library for hatena bookmark collaborative filtering.

## Dependency

- [hbapi](https://github.com/yukihir0/hbapi)
- [hbn](https://github.com/yukihir0/hbn)

## Install

```
go get github.com/yukihir0/hbcf
```

## How to use

```
client := hbn.NewClient()
client.SetTotalPages(2)
client.SetMaxParallelRequest(10)

user := "yukihir0"
bookmarks := client.RequestBookmarks(user)

neighbors := client.SearchNeighbors(bookmarks)
excluded := neighbors.Exclude([]string{
  user,
})
topNeighbors := excluded.Top(20)

c := hbcf.NewClient()
c.SetTotalPages(2)
c.SetMaxParallelRequest(10)

users := hbcf.NewUsers(topNeighbors)
userItems := c.RequestUserItems(users)

strategy := hbcf.NewSimilarityStrategy()
items := hbcf.Recommend(strategy, userItems)
topItems := items.Top(20)

for _, item := range topItems {
  fmt.Printf("- %s : %.3f\n", item.Title, item.Score)
}
```

## License

Copyright &copy; 2015 yukihir0
