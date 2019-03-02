package controller

import (
	"doge-apis/apis"

	routing "github.com/qiangxue/fasthttp-routing"
)

func TestMapStringInterfaceCtrl(c *routing.Context) error {
	c.SetStatusCode(200)
	c.SetBodyString(apis.TestMapStringInterface())
	return nil
}
