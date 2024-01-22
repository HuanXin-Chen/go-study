package model

import (
	"fmt"
)

// 结构体依赖关系

type Message string

type Channel struct {
	Message Message
}

type BroadCast struct {
	Channel Channel
}

type IBroadCast interface {
	Start()
}

func NewMessage() Message {
	return Message("hello go-wire")
}

func NewMessageWithClose() (Message, func(), error) {
	return Message("hello go-wire"), func() {
		fmt.Println("close")
	}, nil
}

func NewChannel(m Message) Channel {
	return Channel{Message: m}
}

func NewBroadcast(c Channel) BroadCast {
	return BroadCast{Channel: c}
}

func (c Channel) GetMsg() Message {
	return c.Message
}

func (b BroadCast) Start() {
	msg := b.Channel.GetMsg()
	fmt.Println(msg)
}

func Single(c Channel) *BroadCast {
	return &BroadCast{Channel: c}
}
