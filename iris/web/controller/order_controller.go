package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris/repositories"
	"iris/service"
)

type Res struct {
	OrderId string `json:"order_id"`
}
type OrderController struct {
}

func (oc *OrderController) Get() mvc.View {
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo("220212444599762")
	orderId := orderInfo.OrderId
	return mvc.View{
		Name: "hello.html",
		Data: orderId,
	}

}
func (oc *OrderController) GetOrderInfo(c iris.Context) {
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo("220212444599762")
	res := Res{
		OrderId: orderInfo.OrderId,
	}
	// 使用转换后的 JSON 字符串作为参数传递给 iris.Context.JSON 方法
	c.JSON(res)
}
