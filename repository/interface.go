package model

import (
	"alta-test/entities"
	"alta-test/view"
)

type ModelUser interface {
	CreateUser(newUser entities.User) (entities.User, error)
	GetAllUsers(name, role, sort string, sizePage, page int) ([]entities.User, error)
	GetUserID(id uint, role string) (entities.User, error)
	UpdateUserID(id uint, update entities.User) (entities.User, error)
	DeleteUserID(id uint) error
	GetUserLogin(login view.Login) (entities.User, error)
}
