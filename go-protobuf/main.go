package main

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"go-protobuf/model"
	"time"
)

func main() {
	t := 1000000
	p := pro(t)
	j := js(t)
	logrus.Infof("times:%f", float64(j)/float64(p))
}

func pro(t int) int64 {
	meg := &model.Message{
		SenderId:   1,
		ReceiverId: 231,
		Data:       []byte("hello"),
		IsGroup:    false,
		Type:       model.Message_Group,
	}
	data := &model.Message{}
	start := time.Now()
	for i := 0; i < t; i++ {
		RawData, _ := proto.Marshal(meg)
		//logrus.Info(RawData)
		proto.Unmarshal(RawData, data)
		//logrus.Info(data)
	}
	end := time.Now()
	logrus.Infof("proto: start:%d, end:%d period:%d",
		start.UnixNano(), end.UnixNano(), end.UnixNano()-start.UnixNano())
	return end.UnixNano() - start.UnixNano()
}

func js(t int) int64 {
	meg := &model.Message{
		SenderId:   1,
		ReceiverId: 231,
		Data:       []byte("hello"),
		IsGroup:    false,
		Type:       model.Message_Group,
	}
	data := &model.Message{}
	start := time.Now()
	for i := 0; i < t; i++ {
		RawData, _ := json.Marshal(meg)
		//logrus.Info(RawData)s
		json.Unmarshal(RawData, data)
		//logrus.Info(data)
	}
	end := time.Now()
	logrus.Infof("json: start:%d, end:%d period:%d",
		start.UnixNano(), end.UnixNano(), end.UnixNano()-start.UnixNano())
	return end.UnixNano() - start.UnixNano()
}
