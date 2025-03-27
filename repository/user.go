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

func (m *UserRepository) Update(id uint, post model.RequestUser) model.ResponseUser {
	_, err := m.Db.Exec("UPDATE users SET username = $1 WHERE id = $2", post.Username, id)
	if err != nil {
		log.Println(err)
		return model.ResponseUser{}
	}

	return m.findOne(id)
}

func (m *UserRepository) findOne(id uint) model.ResponseUser {
	query, err := m.Db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return model.ResponseUser{}
	}

	var dataUser model.ResponseUser

	if query != nil {
		for query.Next() {
			var (
				id       int64
				username string
			)
			err := query.Scan(&id, &username)
			if err != nil {
				log.Println(err)
			}

			dataUser = model.ResponseUser{Id: id, Username: username}
		}
	}

	return dataUser
}

func (m *UserRepository) GetAll() []model.ResponseUser {
	query, err := m.Db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
		return nil
	}

	var listUser []model.ResponseUser

	if query != nil {
		for query.Next() {
			var (
				id       int64
				username string
			)

			err := query.Scan(&id, &username)
			if err != nil {
				log.Println(err)
			}

			dataUser := model.ResponseUser{Id: id, Username: username}
			listUser = append(listUser, dataUser)
		}
	}

	return listUser
}

func (m *UserRepository) Delete(id uint) bool {
	_, err := m.Db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
