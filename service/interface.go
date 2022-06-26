package service

import "alta-test/view"

type ServiceUser interface {
	CreateUser(newUser view.AddUser) (view.RespondUser, error)
	GetAllUsers(name, role, sort string, sizePage, page int) ([]view.RespondUser, error)
	GetUserID(id uint, role string) (view.RespondUser, error)
	UpdateUserID(id uint, update view.UpdateUser) (view.RespondUser, error)
	DeleteUserID(id uint) error
	GetUserLogin(login view.Login) (view.RespondUser, error)
}
