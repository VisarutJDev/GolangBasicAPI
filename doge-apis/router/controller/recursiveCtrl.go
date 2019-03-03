package controller

import (
	"doge-apis/apis"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
)

func TestRecursivetrl(c *routing.Context) error {
	c.SetStatusCode(200)
	apis.TestRecursive()
	response := `{
		"message":  "` + fmt.Sprintf("%s", "test") + `"
	}
	`
	c.SetBodyString(response)
	return nil
}
