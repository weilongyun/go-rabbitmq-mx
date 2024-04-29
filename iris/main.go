package main

import "github.com/kataras/iris/v12"

func main() {
	//初始化iris
	app := iris.New()
	//加载日志等级
	app.Logger().SetLevel("debug")
	//加载前端模板
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//加载控制器
	app.Run(
		iris.Addr(":6789"),
	)
}
