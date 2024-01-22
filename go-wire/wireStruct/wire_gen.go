// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireStruct

import (
	"go-study/go-wire/model"
)

// Injectors from wire.go:

func InitializeBroadCast() model.BroadCast {
	message := model.NewMessage()
	channel := model.NewChannel(message)
	broadCast := model.BroadCast{
		Channel: channel,
	}
	return broadCast
}
