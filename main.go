package main

import (
	"sparrow_rest_framework/apptest"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// app := newApp()
	// 初始化 settings
	// settings.Init()
	// xx := settings.GetString("XX")
	// print(xx)
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	apptest.Register(app, "/api")

	app.Listen("0.0.0.0:8001",
		iris.WithTimeFormat("2006-01-02T15:04:05"),
	)

}
