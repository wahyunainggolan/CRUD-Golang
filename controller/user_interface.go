package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	Create(*gin.Context)
}
