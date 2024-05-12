package main

import (
	"context"
	"log"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	//初始化iris
	app := iris.New()

	//加载日志等级
	app.Logger().SetLevel("debug")
	//加载注册前端模板
	template := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//注册模板目标
	app.HandleDir("/assets", iris.Dir("./backend/web/assets"))
	//c := &controller.OrderController{}
	//注册路由
	//注意：此时只要controller中定义了多少方法都会被执行
	//iris路由有二种模式：mvc路由和函数路由
	//mvc路由模式
	//mvc.New(app.Party("/hello")).Handle(c)
	//函数路由模式，Party代表分组
	/*api := app.Party("/order-group")
	{
	}*/
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
