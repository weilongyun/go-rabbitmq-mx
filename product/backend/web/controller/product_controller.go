package controller

import (
	"backend/common"
	"backend/datamodels"
	"backend/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
)

type Product struct {
	Ctx            iris.Context
	productService service.IProductService
}

func (p *Product) GetAllProductInfo() mvc.View {
	result, err := p.productService.GetAllProductInfo()
	if err != nil {
		log.Fatalln("product controller  GetAllProductInfo error", err)
	}
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"arrAllProductInfo": result,
		},
	}
}

//post请求
func (p *Product) PostUpdateProductInfo() {
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
}
