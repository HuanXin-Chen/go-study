// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireBind

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

// Injectors from wire.go:

func InitializeBroadCast() model.IBroadCast {
	message := model.NewMessage()
	channel := model.NewChannel(message)
	broadCast := model.Single(channel)
	return broadCast
}

// wire.go:

var Bind = wire.Bind(new(model.IBroadCast), new(*model.BroadCast))
