package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"iris/repositories"
	"iris/service"
)

type OrderForm struct {
	OrderId string `form:"order_id"`
}
type OrderQuery struct {
	OrderId string `form:"order_id"`
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
//结构体不需要加标签，校验字段的时候用iris.IsErrPath就会全部校验
func (oc *OrderController) GetOrderInfoByPost(c iris.Context) {
	var form OrderForm
	if err := c.ReadForm(&form); !iris.IsErrPath(err) {
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

//通过ReadQuery绑定参数 适用于get请求
func (oc *OrderController) GetOrderInfoByGet(c iris.Context) {
	var query OrderQuery
	if err := c.ReadQuery(&query); !iris.IsErrPath(err) {
		c.StatusCode(iris.StatusBadRequest)
		logrus.Error("GetOrderInfoByPost error", err)
		return
	}
	order_id := query.OrderId
	repo := repositories.NewOrderRepositoryManager()
	orderService := service.NewOrderServiceManager(repo)
	orderInfo := orderService.ShowOrderInfo(order_id)
	res := Res{
		OrderId: orderInfo.OrderId,
	}
	// 使用转换后的 JSON 字符串作为参数传递给 iris.Context.JSON 方法
	c.JSON(res)

}
