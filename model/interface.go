package model

import (
	"alta-test/entities"
)

type ModelUser interface {
	CreateUser(newUser entities.User) (entities.User, error)
	GetAllUsers(name, role, sort string, sizePage, page int) ([]entities.User, error)
	GetUserID(id uint) (entities.User, error)
	UpdateUserID(id uint, update entities.User) (entities.User, error)
	DeleteUserID(id uint) error
}
