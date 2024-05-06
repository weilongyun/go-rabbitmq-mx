package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/web/controller"
	"log"
	"time"
)

func main() {
	//初始化iris
	app := iris.Default()
	//加载日志等级
	app.Logger().SetLevel("debug")
	//加载前端模板
	app.RegisterView(iris.HTML("./web/views", ".html"))
	c := &controller.OrderController{}
	//注册路由
	//注意：此时只要controller中定义了多少方法都会被执行
	//iris路由有二种模式：mvc路由和函数路由
	//mvc路由模式
	mvc.New(app.Party("/hello")).Handle(c)
	//函数路由模式，Party代表分组
	api := app.Party("/order-group")
	{
		api.Get("/getOrderInfo", c.GetOrderInfo)
		api.Post("/getOrderInfoByPost", c.GetOrderInfoByPost)
		api.Get("/getOrderInfoByGet", c.GetOrderInfoByGet)
		api.Get("/jsonp", c.Jsonp)
	}
	//加载控制器 run是阻塞到
	/*app.Run(
		iris.Addr(":6789"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)*/
	//优雅退出
	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		log.Println("close...")
		// close all hosts.
		app.Shutdown(ctx)
		close(idleConnsClosed)
	})
	go func() {
		lis := app.Listen(":6789", iris.WithOptimizations, iris.WithoutInterruptHandler, iris.WithoutServerError(iris.ErrServerClosed))
		if lis != nil {
			log.Fatalf("Server failed to start: %v", lis)
		}
	}()
	<-idleConnsClosed
}
