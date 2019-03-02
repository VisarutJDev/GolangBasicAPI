package routes

import (
	"doge-apis/router/controller"

	routing "github.com/qiangxue/fasthttp-routing"
)

func RoutePath(r *routing.Router) {
	v := r.Group("/doge")
	v.Get("", controller.HelloThereCtrl)
	v.Get("/speed", controller.TestReceiver)
	v.Get("/all-speed", controller.TestInterface)
	v.Get("/mapstringinterface", controller.TestMapStringInterfaceCtrl)
}
