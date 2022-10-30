package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
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
	SenderId int64
	Data     string
}

func main() {
	eng := gin.Default()
	eng.GET("/ws", handle)
	eng.Run(":7787")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type info struct {
	UserId   int64
	Password string
}

func handle(c *gin.Context) {
	i := info{}
	err := c.ShouldBindJSON(&i)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)

	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
	}
	logrus.Info("Connected")
	for {
		// wsConn.SetReadDeadline(time.Now().Add(10 * time.Minute))
		// _, data, err := wsConn.ReadMessage()
		// if err != nil {
		// 	panic(err)
		// }
		// logrus.Info("read")
		// logrus.Infof("From Client:%+v \n", string(data))
		// //wsConn.WriteMessage(websocket.BinaryMessage, []byte("Server Received:"+string(data)))
		pak := Package{}
		wsConn.ReadJSON(&pak)
		// err = json.Unmarshal(data, &pak)
		// if err != nil {
		// 	panic(err)
		// }
		//logrus.Info(pak)
		if pak.Type == 1 {
			//message
			meg := Input{}
			tmp, err := json.Marshal(pak.Data)
			json.Unmarshal(tmp, &meg)
			if err != nil {
				panic(err)
			}
			logrus.Info(meg)
		} else if pak.Type == 2 {
			//heart
			h := heart{}
			tmp, err := json.Marshal(pak.Data)
			json.Unmarshal(tmp, &h)
			if err != nil {
				panic(err)
			}
			logrus.Info(h)
		}
	}
}
