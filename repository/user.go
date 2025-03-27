package repository

import (
	"database/sql"
	"log"

	"crud/model"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{Db: db}
}

func (m *UserRepository) Create(post model.RequestUser) model.ResponseUser {
	query, err := m.Db.Query("SELECT * FROM users WHERE username = $1", post.Username)
	if err != nil {
		log.Println(err)
	}

	var result model.ResponseUser
	var count int

	for query.Next() {
		count++
	}

	if count == 0 {
		dataUser, err := m.Db.Exec(
			"INSERT INTO users(username) VALUES ($1)",
			post.Username,
		)
		if err != nil {
			log.Println(err)
		}

		idUser, _ := dataUser.LastInsertId()
		result = model.ResponseUser{Id: idUser, Username: post.Username}
	} else {
		return model.ResponseUser{}
	}

	return result
}
