package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/web/controller"
)

func main() {
	//初始化iris
	app := iris.New()
	//加载日志等级
	app.Logger().SetLevel("debug")
	//加载前端模板
	app.RegisterView(iris.HTML("./web/views", ".html"))
	c := &controller.OrderController{}
	//注册路由
	//注意：此时只要controller中定义了多少方法都会被执行
	mvc.New(app.Party("/hello")).Handle(c)
	app.Get("/getOrderInfo", c.GetOrderInfo)
	//加载控制器
	app.Run(
		iris.Addr(":6789"),
	)
}
