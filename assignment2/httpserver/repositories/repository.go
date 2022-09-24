package repositories

import "assignment2/httpserver/repositories/models"

type OrderRepo interface {
	CreateOrder(student *models.Order) error
	UpdateOrder(student *models.Order) error
	FindAllOrder() ([]models.Order, error)
	DeleteOrder(id int) error
}

type ItemRepo interface {
	CreateItem(item *models.Item) error
	CreateManyItems(items []models.Item) error
	UpdateItem(item *models.Item) error
	FindAllItemsByOrderId(orderId int) ([]models.Item, error)
	DeleteAllItemByOrderId(orderId int) error
}
