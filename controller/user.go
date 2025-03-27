package controller

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"crud/model"
	"crud/repository"
)

type UserController struct {
	Db *sql.DB
}

func NewUserController(db *sql.DB) UserControllerInterface {
	return &UserController{Db: db}
}

func (m *UserController) Create(c *gin.Context) {
	var post model.RequestUser
	DB := m.Db

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "message": err})

		return
	}

	repository := repository.NewUserRepository(DB)
	dataUser := repository.Create(post)

	if (dataUser != model.ResponseUser{}) {
		c.JSON(200, gin.H{"user_id": dataUser.Id, "username": dataUser.Username})

		return
	} else {
		c.JSON(200, gin.H{"status": "success", "message": "Data User Already Exists"})

		return
	}
}
