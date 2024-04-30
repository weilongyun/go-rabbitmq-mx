package repositories

import "iris/datamodels"

type OrderRepository interface {
	GetOrderInfo(order_id string) *datamodels.OrderModel
}
type OrderRepositoryManager struct {
}

func NewOrderRepositoryManager() OrderRepository {
	return &OrderRepositoryManager{}
}

func (o *OrderRepositoryManager) GetOrderInfo(order_id string) *datamodels.OrderModel {
	//模拟链接数据库，实际上没链接数据库，这里只是定义各个层级的规范
	return &datamodels.OrderModel{
		Id: order_id,
	}
}
