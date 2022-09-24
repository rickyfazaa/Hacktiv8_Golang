package params

import (
	"assignment2/httpserver/repositories/models"
)

type CreateItemRequest struct {
	ItemCode    string `json:"itemCode" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

func (r *CreateItemRequest) ParseToModel() *models.Item {
	return &models.Item{
		ItemCode:    r.ItemCode,
		Description: r.Description,
		Quantity:    r.Quantity,
	}
}

type UpdateItemRequest struct {
	ItemId      int    `json:"itemId" validate:"required"`
	ItemCode    string `json:"itemCode" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

func (r *UpdateItemRequest) ParseToModel() *models.Item {
	return &models.Item{
		ItemId:      r.ItemId,
		ItemCode:    r.ItemCode,
		Description: r.Description,
		Quantity:    r.Quantity,
	}
}
