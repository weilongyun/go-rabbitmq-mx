package controller

import (
	"backend/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ProductController struct {
	//Ctx            iris.Context
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
