package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetAll(*gin.Context)
}
