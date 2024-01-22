//go:build wireinject
// +build wireinject

package wireStruct

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

func InitializeBroadCast() model.BroadCast {
	wire.Build(
		wire.Struct(new(model.BroadCast), "*"),
		model.NewChannel,
		model.NewMessage,
	)
	return model.BroadCast{}
}
