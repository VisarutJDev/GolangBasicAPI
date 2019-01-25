package controller

import (
	"doge-apis/apis"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func TestReceiver(c *routing.Context) error {
	c.SetStatusCode(200)
	response := `{
		"message":  "` + fmt.Sprintf("%s", apis.TestReceiver()) + `"
	}
	`
	c.SetBodyString(response)
	return nil
}
