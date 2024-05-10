package datamodels

//product模型
type Product struct {
	ID           int64  `json:"product_id" sql:"product_id"`
	ProductName  string `json:"product_name" sql:"product_name"`
	ProductNum   int64  `json:"product_num" sql:"product_num"`
	ProductImage string `json:"image" sql:"image"`
	ProductPrice string `json:"product_price" sql:"product_price"`
}
