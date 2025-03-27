package repository

import "crud/model"

type UserRepositoryInterface interface {
	Create(model.RequestUser) model.ResponseUser
	Update(uint, model.RequestUser) model.ResponseUser
	Delete(uint) bool
	GetAll() []model.ResponseUser
}
