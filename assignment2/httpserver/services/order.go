package services

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/controllers/views"
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/repositories/models"
	"net/http"
)

type OrderSvc struct {
	repo     repositories.OrderRepo
	itemRepo repositories.ItemRepo
}

func NewOrderSvc(repo repositories.OrderRepo, itemRepo repositories.ItemRepo) *OrderSvc {
	return &OrderSvc{
		repo:     repo,
		itemRepo: itemRepo,
	}
}

func (s *OrderSvc) CreateOrder(req *params.CreateOrderRequest) *views.Response {
	order, err := req.ParseToModel()
	if err != nil {
		return views.ErrorResp(http.StatusBadRequest, views.BAD_REQUEST, err)
	}

	err = s.repo.CreateOrder(order)
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	items := make([]models.Item, 0)
	for _, item := range req.Items {
		mItem := *item.ParseToModel()
		mItem.OrderId = order.OrderId
		items = append(items, mItem)
	}

	err = s.itemRepo.CreateManyItems(items)
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResp(http.StatusCreated, views.CREATED, nil)
}

func (s *OrderSvc) UpdateOrder(req *params.UpdateOrderRequest) *views.Response {
	order, err := req.ParseToModel()
	if err != nil {
		return views.ErrorResp(http.StatusBadRequest, views.BAD_REQUEST, err)
	}

	err = s.repo.UpdateOrder(order)
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	for _, item := range req.Items {
		mItem := item.ParseToModel()
		mItem.OrderId = order.OrderId

		err = s.itemRepo.UpdateItem(mItem)
		if err != nil {
			return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
		}
	}

	return views.SuccessResp(http.StatusOK, views.OK, nil)
}

func (s *OrderSvc) FindAllOrder() *views.Response {
	orders, err := s.repo.FindAllOrder()
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	vOrders := make([]views.OrderGetAll, 0)
	for _, order := range orders {
		vOrder := views.OrderGetAll{
			OrderId:      order.OrderId,
			OrderedAt:    order.OrderedAt.String(),
			CustomerName: order.CustomerName,
		}

		items, err := s.itemRepo.FindAllItemsByOrderId(order.OrderId)
		if err != nil {
			return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
		}

		vItems := make([]views.ItemGetAll, 0)
		for _, item := range items {
			vItem := views.ItemGetAll{
				ItemId:      item.ItemId,
				ItemCode:    item.ItemCode,
				Description: item.Description,
				Quantity:    item.Quantity,
			}

			vItems = append(vItems, vItem)
		}
		vOrder.Items = vItems
		vOrders = append(vOrders, vOrder)
	}
	return views.SuccessResp(http.StatusOK, views.OK, vOrders)
}

func (s *OrderSvc) DeleteOrder(id int) *views.Response {
	err := s.repo.DeleteOrder(id)
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	err = s.itemRepo.DeleteAllItemByOrderId(id)
	if err != nil {
		return views.ErrorResp(http.StatusInternalServerError, views.INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResp(http.StatusOK, views.OK, nil)
}
