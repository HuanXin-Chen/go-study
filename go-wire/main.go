package main

import (
	"go-study/go-wire/wireClean"
)

func main() {
	b, clean, _ := wireClean.InitializeBroadCast()
	b.Start()
	clean()
}
