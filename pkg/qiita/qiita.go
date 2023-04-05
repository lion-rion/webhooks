package qiita

import (
	. "apihandler/pkg/discord"
	. "apihandler/pkg/fetch"
	"encoding/json"
	"fmt"
)

type Node struct {
	Url        string `json:"linkUrl"`
	LikesCount int    `json:"likesCount"`
	Title      string `json:"title"`
}

type Item struct {
	Node Node `json:"node"`
}

func GetQiitaTrend(whurl string) {
	name := "Qiita"
	url := "https://qiita-api.vercel.app/api/trend"
	img := "https://pbs.twimg.com/profile_images/1610819875567734785/5kM_BxFL_400x400.jpg"

	body := Fetch(url)

	var jdata []Item
	if err := json.Unmarshal([]byte(body), &jdata); err != nil {
		fmt.Println(err)
		return
	}

	message := "🌟 本日の人気記事です 🌟\n\n"

	for i := 0; i < 5; i++ {
		message += fmt.Sprintf("%s ❤️%d \n%s \n\n", jdata[i].Node.Title, jdata[i].Node.LikesCount, jdata[i].Node.Url)
	}

	SendWebhook(name, whurl, img, message)

}
