package main

import (
	path "doge-apis/router/routes"
	"fmt"

	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	r := routing.New()
	r.Use(errorHandler)
	r.NotFound(notFoundHandler)

	path.RoutePath(r)

	h := r.HandleRequest
	h = fasthttp.CompressHandler(h)

	panic(fasthttp.ListenAndServe(":8080", h))
}

func notFoundHandler(c *routing.Context) error {
	c.RequestCtx.NotFound()
	c.Next()
	return nil
}

func errorHandler(c *routing.Context) error {
	c.SetContentType("application/json")
	defer func(c *routing.Context) {
		if rec := recover(); rec != nil {
			fmt.Println("Recovered in f", rec)
			switch x := rec.(type) {
			case string:
				response := `{
					"message": "` + fmt.Sprintf("%s", x) + `"
				}`
				c.SetStatusCode(400)
				c.RequestCtx.SetBodyString(response)
			case error:
				response := `{
					"message": "` + fmt.Sprintf("%s", x.Error()) + `"
				}`
				c.SetStatusCode(400)
				c.RequestCtx.SetBodyString(response)
			default:
				response := `{
					"message": "Something wrong. Fix it baka!"
				}`
				c.SetStatusCode(400)
				c.RequestCtx.SetBodyString(response)
			}
			c.Abort()
		}
	}(c)
	c.Next()
	return nil
}
