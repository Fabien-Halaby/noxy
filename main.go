package main

import (
	"noxy/controller"
	"noxy/model"
	"noxy/view"
)

func main() {
	cache := model.NewCache()
	args := controller.ParseArgs()

	if args.ClearCache {
		cache.Clear()
		view.ShowClearMessage()

		return
	}

	controller.StartServer(cache, args)
}
