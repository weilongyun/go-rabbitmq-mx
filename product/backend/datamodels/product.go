package datamodels

//product模型
type Product struct {
	ID           int64  `json:"id" sql:"id" mx:"id"`
	ProductID    string `json:"product_id" sql:"product_id" mx:"product_id"`
	ProductName  string `json:"product_name" sql:"product_name" mx:"product_name"`
	ProductNum   int    `json:"product_num" sql:"product_num" mx:"product_num"`
	ProductImage string `json:"image" sql:"image" mx:"product_image"`
	ProductPrice int    `json:"product_price" sql:"product_price" mx:"product_price"`
}
