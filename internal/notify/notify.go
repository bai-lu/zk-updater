package notify

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Message map[string]string

var queue = make(chan Message, 100)

func init() {
	go run()
}

// 把消息推入队列
func Push(msg Message) {
	queue <- msg
}

func run() {
	for msg := range queue {
		// 根据任务配置发送通知
		_, receiverOk := msg["receivers"]
		_, titleOk := msg["title"]
		_, textOk := msg["text"]
		_, imgOk := msg["img"]
		if !receiverOk || !titleOk || !textOk || !imgOk {
			log.Errorf("#notify#参数不完整#%+v", msg)
			continue
		}

		// 飞书
		lark := Lark{}
		go lark.Send(msg)
		time.Sleep(1 * time.Second)
	}
}
