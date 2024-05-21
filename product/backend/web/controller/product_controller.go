package controller

import (
	"backend/common"
	"backend/datamodels"
	"backend/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
/*func (p *Product) PostUpdateProductInfo() {
	product := &datamodels.Product{}
	//解析表单
	p.Ctx.Request().ParseForm()
	decoder := common.NewDecoder(&common.DecoderOptions{})
	if err := decoder.Decode(p.Ctx.Request().Form, product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err := p.productService.UpdateProduct(product)
	if err != nil {
		log.Fatalln("product controller  GetAllProductInfo error", err)
	}
}*/
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
