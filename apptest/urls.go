package apptest

import (
	// "github.com/kataras/iris//v12/mvc"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Register(app *iris.Application, apiPrefix string) {
	appParty := app.Party(apiPrefix)
	{
		mApp := mvc.New(appParty)
		// mApp.Router.Use(auth.LoginRequire)
		c := NewController()
		mApp.Handle(c)
	}
}
