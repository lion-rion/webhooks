package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type discordWebhook struct {
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	Content   string `json:"content"`
}

func SendWebhook(name string, whurl string, imgurl string, content string) {

	message := discordWebhook{UserName: name, AvatarURL: imgurl, Content: content}

	j, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}

	req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
	if err != nil {
		fmt.Println("new request err:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	if res.StatusCode == 204 {
		fmt.Println("success") //成功
	} else {
		fmt.Printf("### ERROR ###\n\n%#v\n", res) //失敗
	}

}
