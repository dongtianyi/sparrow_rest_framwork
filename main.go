package main

import (
	"sparrow_rest_framework/apptest"
	"sparrow_rest_framework/database"
	"sparrow_rest_framework/settings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// 初始化 settings
	settings.Init()
	dbUser := settings.GetString("DATABASE_USER")
	dbPwd := settings.GetString("DATABASE_PASSWORD")
	dbHost := settings.GetString("DATABASE_HOST")
	dbPort := settings.GetString("DATABASE_PORT")
	dbName := settings.GetString("DATABASE_DBNAME")
	database.InitDb(dbUser, dbPwd, dbHost, dbPort, dbName)

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	apptest.Register(app, "/api")

	app.Listen("0.0.0.0:8001",
		iris.WithTimeFormat("2006-01-02T15:04:05"),
	)

}
