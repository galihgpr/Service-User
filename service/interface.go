package service

import "alta-test/view"

type ServiceUser interface {
	CreateUser(newUser view.AddUser) (view.RespondUser, error)
	GetAllUsers() ([]view.RespondUser, error)
	GetUserID(id uint) (view.RespondUser, error)
	UpdateUser(id uint, update view.UpdateUser) (view.RespondUser, error)
	DeleteUserID(id uint) error
}