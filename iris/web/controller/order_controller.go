package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"iris/repositories"
	"iris/service"
	"os"
)

type OrderForm struct {
	OrderId string `form:"order_id" binding:"required"`
}
type OrderQuery struct {
	OrderId string `form:"order_id" binding:"required"`
}
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

//http://localhost:6789/getOrderInfo?order_id=11
func (oc *OrderController) GetOrderInfo(c iris.Context) {
	order_id := c.URLParam("order_id")
	if len(order_id) == 0 {
		logrus.Error("order_id error", order_id)
		return
	}
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo(order_id)
	res := Res{
		OrderId: orderInfo.OrderId,
	}
	// 使用转换后的 JSON 字符串作为参数传递给 iris.Context.JSON 方法
	c.JSON(res)
}

//通过ReadForm绑定参数 适用于post请求
func (oc *OrderController) GetOrderInfoByPost(c iris.Context) {
	var form OrderForm
	//这个绑定没有用，只是为了方便获取字段值的，后续强制校验的字段还需要业务逻辑去判断，没有gin灵活
	if err := c.ReadForm(&form); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		logrus.Error("GetOrderInfoByPost error", err)
		return
	}
	order_id := form.OrderId
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo(order_id)
	res := Res{
		OrderId: orderInfo.OrderId,
	}
	// 使用转换后的 JSON 字符串作为参数传递给 iris.Context.JSON 方法
	c.JSON(res)
}

//通过ReadForm绑定参数 适用于post请求
func (oc *OrderController) GetOrderInfoByGet(c iris.Context) {
	var form OrderQuery
	if err := c.ReadQuery(&form); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		logrus.Error("GetOrderInfoByPost error", err)
		return
	}
	order_id := form.OrderId
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo(order_id)
	res := Res{
		OrderId: orderInfo.OrderId,
	}
	// 使用转换后的 JSON 字符串作为参数传递给 iris.Context.JSON 方法
	c.JSON(res)

}
