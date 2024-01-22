//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

func InitializeBroadCast() model.BroadCast {
	wire.Build(model.NewBroadcast, model.NewChannel, model.NewMessage)
	return model.BroadCast{}
}
