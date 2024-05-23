package controller

import (
	"backend/common"
	"backend/datamodels"
	"backend/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"strconv"
)

type ProductController struct {
	Ctx            iris.Context
	ProductService service.IProductService
}

//获取全部商品列表
func (p *ProductController) GetAll() mvc.View {
	productArray, _ := p.ProductService.GetAllProductInfo()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}

//post请求
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	//解析表单
	p.Ctx.Request().ParseForm()
	p.Ctx.Application().Logger().Info("ProductController PostUpdate start")
	decoder := common.NewDecoder(&common.DecoderOptions{})
	if err := decoder.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		log.Fatalln("product controller  GetAllProductInfo error", err)
	}
	p.Ctx.Redirect("/product/all")
}

//添加商品表单页
func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name: "product/add.html",
	}
}

//添加商品
func (p *ProductController) PostAdd() {
	product := &datamodels.Product{}
	//解析post字段到结构体中
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{})
	if err := dec.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProdct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

//获取商品详情
func (p *ProductController) GetManager() mvc.View {
	idStr := p.Ctx.URLParam("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		p.Ctx.Application().Logger().Error("GetManager Atoi error", err)
		return mvc.View{
			Code: iris.StatusBadRequest, // 设置响应状态码为 400 Bad Request
			Name: "error.html",          // 指定渲染的视图模板
			Data: iris.Map{ // 传递给视图模板的数据
				"Error": "Invalid ID", // 错误消息
			},
		}
	}
	resp, err := p.ProductService.GetProductInfoByID(int64(idInt))
	if err != nil {
		p.Ctx.Application().Logger().Error("GetManager GetProductInfoByID  error", err)
		return mvc.View{
			Code: iris.StatusBadRequest, // 设置响应状态码为 400 Bad Request
			Name: "error.html",          // 指定渲染的视图模板
			Data: iris.Map{ // 传递给视图模板的数据
				"Error": "Invalid ID", // 错误消息
			},
		}
	}
	return mvc.View{
		Name: "product/manager.html",
		Data: iris.Map{
			"product": resp,
		},
	}
}
func (p *ProductController) GetDelete() {
	idStr := p.Ctx.URLParam("id")
	p.Ctx.Application().Logger().Info("GetDelete  idStr is", idStr)
	if len(idStr) == 0 {
		p.Ctx.Application().Logger().Info("GetDelete  idStr empty")
		return
	}
	resp, err := p.ProductService.DeletePruductByID(idStr)
	p.Ctx.Application().Logger().Info("GetDelete  DeletePruductByID resp ", resp)
	if err != nil {
		p.Ctx.Application().Logger().Error("GetDelete  DeletePruductByID error", err)
		return
	}
	p.Ctx.Redirect("/product/all")

}
