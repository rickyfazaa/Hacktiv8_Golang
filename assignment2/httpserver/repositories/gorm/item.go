package gorm

import (
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"

	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) repositories.ItemRepo {
	return &itemRepo{
		db: db,
	}
}

func (r *itemRepo) CreateItem(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepo) CreateManyItems(items []models.Item) error {
	return r.db.Create(&items).Error
}

func (r *itemRepo) UpdateItem(item *models.Item) error {
	return r.db.Model(item).Updates(*item).Error
}

func (r *itemRepo) FindAllItemsByOrderId(orderId int) ([]models.Item, error) {
	items := make([]models.Item, 0)
	result := r.db.Where("order_id = ?", orderId).Find(&items)
	return items, result.Error
}

func (r *itemRepo) DeleteAllItemByOrderId(orderId int) error {
	return r.db.Delete(&models.Item{}, "order_id = ?", orderId).Error
}
