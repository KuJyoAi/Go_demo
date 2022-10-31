package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type BotMessage struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
	Type    int64  `json:"type"`
}

func main() {
	meg := &BotMessage{
		Sender:  "111",
		Content: "222",
		Type:    0,
	}
	s, err := json.Marshal(meg)
	if err != nil {
		fmt.Println(err)
	}
	Post("http://127.0.0.1:8080/", s)

}

func SendToDiscordBot(meg *BotMessage) {
	s, err := json.Marshal(meg)
	if err != nil {
		fmt.Println(err)
	}
	Post("http://127.0.0.1:5702/", s)
}

// Get 发送GET请求
func Get(url string) error {
	// 超时时间：1分钟
	client := &http.Client{Timeout: 1 * time.Minute}
	resp, err := client.Get(url)
	if err != nil {
		logrus.Errorf("[util.get]err=%+v,url=%+v", err, url)
		return err
	}
	resp.Body.Close()
	return nil
}

// Post 发送Post请求
func Post(url string, body []byte) error {
	client := &http.Client{Timeout: 1 * time.Minute}
	r := bytes.NewReader(body)
	resp, err := client.Post(url, "", r)
	if err != nil {
		logrus.Errorf("[util.Post] Post %+v", err)
		return err
	}
	resp.Body.Close()
	return nil
}
