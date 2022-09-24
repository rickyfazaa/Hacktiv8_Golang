package controllers

import (
	"net/http"
	"sesi6-gin/httpserver/controllers/params"
	"sesi6-gin/httpserver/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUser(ctx *gin.Context) {
	var req params.UserCreateRequest
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

	response := services.CreateUser(&req)

	WriteJsonRespnse(ctx, response)
}

func ReadUser(ctx *gin.Context) {
	response := services.ReadUser()
	WriteJsonRespnse(ctx, response)
}
