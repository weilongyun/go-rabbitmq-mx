package service

import (
	"backend/datamodels"
	"backend/repositories"
)

type IProductService interface {
	InsertProdct(*datamodels.Product) (int64, error)
	DeletePruductByProductid(string2 string) (bool, error)
	UpdateProduct(*datamodels.Product) error
	GetProductInfoByProductid(string) (*datamodels.Product, error)
	GetAllProductInfo() ([]*datamodels.Product, error)
}
type ProductServiceManager struct {
	productRepo repositories.IProductRepository
}

//实例化
func NewProductServiceManager(productRepo repositories.IProductRepository) IProductService {
	return &ProductServiceManager{
		productRepo: productRepo,
	}
}
func (p *ProductServiceManager) InsertProdct(product *datamodels.Product) (int64, error) {
	return p.productRepo.Insert(product)
}

func (p *ProductServiceManager) DeletePruductByProductid(product_id string) (bool, error) {
	return p.productRepo.Delete(product_id)
}

func (p *ProductServiceManager) UpdateProduct(product *datamodels.Product) error {
	return p.productRepo.Update(product)
}

func (p *ProductServiceManager) GetProductInfoByProductid(product_id string) (*datamodels.Product, error) {
	return p.productRepo.SelectByProductId(product_id)
}

func (p *ProductServiceManager) GetAllProductInfo() ([]*datamodels.Product, error) {
	return p.productRepo.SelectAll()
}
