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
	OrderId string `url:"order_id,required"` //get方式用关键字url
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
//IsErrPath只是适用于post，不适用于get
//iris文档 https://www.tizi365.com/topic/10664.html
func (oc *OrderController) GetOrderInfoByPost(c iris.Context) {
	var form OrderForm
	err := c.ReadForm(&form)
	//IsErrPath会忽略不存在的参数，还是会继续向后执行的，所以post方式不能强制校验，需要业务去判断
	if err != nil && !iris.IsErrPath(err) {
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
//http://localhost:6789/getOrderInfoByGet?order_id=1
func (oc *OrderController) GetOrderInfoByGet(c iris.Context) {
	var query OrderQuery
	//get方式可以强制校验参数
	if err := c.ReadQuery(&query); err != nil {
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
