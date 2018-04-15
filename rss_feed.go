package main

import (
	"bytes"
	"fmt"
	"github.com/mmcdole/gofeed"
	"net/http"
	"os"
)

func main() {
	var url string = os.Args[1]
	var webhook string = os.Args[2]

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	for i := 0; i < len(feed.Items); i++ {
		postText := feed.Items[i].Title + "\n" + feed.Items[i].Content + "\n" + feed.Items[i].Link

		jsonStr := "{\"text\":\"" + postText + "\"}"
		fmt.Print(jsonStr)
		req, _ := http.NewRequest(
			"POST",
			webhook,
			bytes.NewBuffer([]byte(jsonStr)),
		)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Print(err)
		}
		defer resp.Body.Close()
	}
}
