package controllers

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderController struct {
	svc *services.OrderSvc
}

func NewOrderController(svc *services.OrderSvc) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.CreateOrderRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.CreateOrder(&req)
	WriteJsonResponse(ctx, response)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var req params.UpdateOrderRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.UpdateOrder(&req)
	WriteJsonResponse(ctx, response)
}

func (c *OrderController) GetAllOrder(ctx *gin.Context) {
	response := c.svc.FindAllOrder()
	WriteJsonResponse(ctx, response)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	response := c.svc.DeleteOrder(id)
	WriteJsonResponse(ctx, response)
}
