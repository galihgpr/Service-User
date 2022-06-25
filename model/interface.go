package model

import (
	"alta-test/entities"
	"alta-test/view"
)

type ModelUser interface {
	CreateUser(newUser entities.User) (entities.User, error)
	GetAllUsers() ([]view.RespondUser, error)
	GetUserID(id uint) (entities.User, error)
	UpdateUserID(id uint, update entities.User) (entities.User, error)
	DeleteUserID(id uint) error
}
