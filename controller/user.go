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

func (m *UserController) Update(c *gin.Context) {
	DB := m.Db
	var post model.RequestUser
	var uri model.UserUri

	if err := c.ShouldBind(&post); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewUserRepository(DB)
	update := repository.Update(uri.ID, post)

	if (update != model.ResponseUser{}) {
		c.JSON(200, gin.H{"status": "success", "data": update, "msg": "update user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "data": nil, "msg": "update user failed"})
		return
	}
}

func (m *UserController) Delete(c *gin.Context) {
	DB := m.Db
	var uri model.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "failed", "msg": err})
		return
	}

	repository := repository.NewUserRepository(DB)
	delete := repository.Delete(uri.ID)

	if delete {
		c.JSON(200, gin.H{"status": "success", "msg": "delete user successfully"})
		return
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "delete user failed"})
		return
	}
}

func (m *UserController) GetAll(c *gin.Context) {
	DB := m.Db
	repository := repository.NewUserRepository(DB)
	get := repository.GetAll()

	if get != nil {
		c.JSON(200, gin.H{"status": "success", "data": get, "msg": "get user successfully"})
		return
	} else {
		c.JSON(200, gin.H{"status": "success", "data": nil, "msg": "user not found"})
		return
	}
}
