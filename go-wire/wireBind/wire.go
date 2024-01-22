//go:build wireinject
// +build wireinject

package wireBind

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

var Bind = wire.Bind(new(model.IBroadCast), new(*model.BroadCast))

func InitializeBroadCast() model.IBroadCast {
	wire.Build(model.NewMessage, model.NewChannel, model.Single, Bind)
	return model.BroadCast{}
}
