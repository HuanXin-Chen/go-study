//go:build wireinject
// +build wireinject

package wireClean

import (
	"github.com/google/wire"
	"go-study/go-wire/model"
)

func InitializeBroadCast() (model.BroadCast, func(), error) {
	wire.Build(model.NewBroadcast, model.NewChannel, model.NewMessageWithClose)
	return model.BroadCast{}, nil, nil
}
