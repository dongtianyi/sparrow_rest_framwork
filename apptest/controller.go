package apptest

import (
	"sparrow_rest_framework/rest_framwork"

	"github.com/kataras/iris/v12"
)

// type IController interface {
// 	GetTest(ctx iris.Context) rest_framwork.ResultData
// }

type Controller struct {
	Ctx   iris.Context
	Model interface{}
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetTest(ctx iris.Context) rest_framwork.ResultData {
	return rest_framwork.ResponseResult(c, 1, "ok")
}

/*
	保留 django 的风格
*/
type GenericAPIView struct {
	Model interface{}
}

type ListAPIView struct {
}

func (v *ListAPIView) Get() {

}
