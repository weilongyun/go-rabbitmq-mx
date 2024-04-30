package service

import "iris/repositories"

type OrderResp struct {
	OrderId string `json:"order_id"`
}
type OrderService interface {
	ShowOrderInfo(order_id string) *OrderResp
}
type OrderServiceManager struct {
	repo repositories.OrderRepository
}

func NewOrderServiceManager(repo repositories.OrderRepository) OrderService {
	return &OrderServiceManager{
		repo: repo,
	}
}

func (o *OrderServiceManager) ShowOrderInfo(order_id string) *OrderResp {
	orderInfo := o.repo.GetOrderInfo(order_id)
	return &OrderResp{
		OrderId: orderInfo.Id,
	}
}
