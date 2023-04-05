package zenn

import (
	. "apihandler/pkg/discord"
	. "apihandler/pkg/fetch"
	"encoding/json"
	"fmt"
)

type Data struct {
	Title string `json:"title"`
	Like  int    `json:"likedCount"`
	Emoji string `json:"emoji"`
	Path  string `json:"path"`
}

func GetZennTrend(whurl string) {
	url := "https://zenn-api.vercel.app/api/trendTech"
	img := "https://gyazo.com/269f519fde2f7a3f4c699e073998f238/max_size/400"

	body := Fetch(url)
	var events []Data
	if err := json.Unmarshal([]byte(body), &events); err != nil {
		fmt.Println(err)
		return
	}

	message := "🌟 本日の人気記事です 🌟\n\n"
	// ここで、eventsの中身をループで5回して、messageに追加していく
	for i := 0; i < 5; i++ {
		message += fmt.Sprintf("%s %s ❤️%d \n https://zenn.dev/%s \n\n\n", events[i].Emoji, events[i].Title, events[i].Like, events[i].Path)
	}

	SendWebhook("zenn", whurl, img, message)

}
