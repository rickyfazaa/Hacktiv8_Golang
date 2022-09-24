package gorm

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}
func (r *orderRepo) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) UpdateOrder(order *models.Order) error {
	return r.db.Model(order).Updates(*order).Error
}

func (r *orderRepo) FindAllOrder() ([]models.Order, error) {
	orders := make([]models.Order, 0)
	result := r.db.Find(&orders)

	return orders, result.Error
}

func (r *orderRepo) DeleteOrder(id int) error {
	return r.db.Delete(&models.Order{}, id).Error
}
