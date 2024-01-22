//go:build wireinject
// +build wireinject

package wireSet

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

var providerSet = wire.NewSet(model.NewBroadcast, model.NewChannel, model.NewMessage)

func InitializeBroadCast() model.BroadCast {
	wire.Build(providerSet)
	return model.BroadCast{}
}
