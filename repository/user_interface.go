package repository

import "crud/model"

type UserRepositoryInterface interface {
	Create(model.RequestUser) model.ResponseUser
}
