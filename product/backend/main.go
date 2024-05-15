package main

import (
	"backend/common"
	"backend/repositories"
	"backend/service"
	"backend/web/controller"
	"context"
	"github.com/kataras/iris/v12/mvc"
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
	productInstance := mvc.New(app.Party("/product"))
	dbConn, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalf("start NewMysqlConn err %v", err)
	}
	productRepositories := repositories.NewProductRepositoryManager("product", dbConn)
	productService := service.NewProductServiceManager(productRepositories)
	//注册service
	productInstance.Register(productService)
	//注册控制器
	productInstance.Handle(new(controller.Product))
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
