package params

import (
	"assignment2/httpserver/repositories/models"
	"time"
)

type CreateOrderRequest struct {
	OrderedAt    string              `json:"orderedAt" validate:"required"`
	CustomerName string              `json:"customerName" validate:"required"`
	Items        []CreateItemRequest `json:"items" validate:"required,dive"`
}

func (r *CreateOrderRequest) ParseToModel() (*models.Order, error) {
	t, err := time.Parse(time.RFC3339, r.OrderedAt)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		OrderedAt:    t,
		CustomerName: r.CustomerName,
	}, nil
}

type UpdateOrderRequest struct {
	OrderId      int                 `json:"orderId" validate:"required"`
	OrderedAt    string              `json:"orderedAt" validate:"required"`
	CustomerName string              `json:"customerName" validate:"required"`
	Items        []UpdateItemRequest `json:"items" validate:"required,dive"`
}

func (r *UpdateOrderRequest) ParseToModel() (*models.Order, error) {
	t, err := time.Parse(time.RFC3339, r.OrderedAt)
	if err != nil {
		return nil, err
	}

	return &models.Order{
		OrderId:      r.OrderId,
		OrderedAt:    t,
		CustomerName: r.CustomerName,
	}, nil
}
