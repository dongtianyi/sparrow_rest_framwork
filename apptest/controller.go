package apptest

import (
	"sparrow_rest_framework/rest_framwork"

	"github.com/kataras/iris/v12"
)

type Controller struct {
	Ctx iris.Context
	v   rest_framwork.ListAPIView
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetTest(ctx iris.Context) rest_framwork.ResultData {
	x := PdtPdt{}
	c.v.SetModel(x)
	result, err := c.v.List()
	return rest_framwork.ResponseResult(result, 0, err)
}
