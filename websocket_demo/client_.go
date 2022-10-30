package main

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type PackageType int
type Package struct {
	//数据包内容, 按需修改
	Type PackageType
	Data interface{}
}

type Input struct {
	SenderId int64  `json:"send_id"`
	GroupId  int64  `json:"group_id"`
	Data     string `json:"data"`
	Heart    heart  `json:"heart"`
}
type heart struct {
	SenderId int64  `json:"send_id"`
	Data     string `json:"data"`
}

func main() {
	//conn, resp, err := websocket.DefaultDialer.Dial("wss://chat.pivotstudio.cn/ws?email=1761373255@qq.com&password=xhzq233", nil)
	conn, resp, err := websocket.DefaultDialer.Dial("ws://127.0.0.1:7787/ws?user_id=1&password=123456", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp != nil {
		logrus.Infof("Server:%s\n", string(bytes))
	}

	//go HeartBeat(conn)
	go Message(conn)
	Receive(conn)
}
func Message(c *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 4)
	h := heart{
		SenderId: 1,
		Data:     "Heart",
	}
	meg := Input{
		SenderId: 1,
		GroupId:  2,
		Data:     "Hello",
		Heart:    h,
	}

	for range ticker.C {

		sh := Package{
			Type: 2,
			Data: h,
		}
		s := Package{
			Type: 1,
			Data: meg,
		}
		c.WriteJSON(sh)
		logrus.Info("心跳发送")
		// OutputData, err := json.Marshal(s)
		// if err != nil {
		// 	panic(err)
		// }
		// c.WriteMessage(websocket.BinaryMessage, OutputData)
		c.WriteJSON(s)
		logrus.Info("消息发送")
	}
}
func HeartBeat(c *websocket.Conn) {
	ticker := time.NewTicker(time.Second * 4)
	h := heart{
		SenderId: 1,
		Data:     "Heart",
	}
	h_json, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	for range ticker.C {
		s := Package{
			Type: 2,
			Data: h_json,
		}
		c.WriteJSON(s)
		logrus.Info("心跳发送")
	}
}

func Receive(c *websocket.Conn) {
	c.SetReadDeadline(time.Now().Add(10 * time.Minute))
	for {
		_, bytes, err := c.ReadMessage()
		if err != nil {
			logrus.Infof("[Receive], %+v", err)
			return
		}
		logrus.Infof("Server:%s\n", string(bytes))
		p := Package{}
		json.Unmarshal(bytes, &p)
		logrus.Info(p)
	}
}
