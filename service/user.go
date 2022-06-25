package service

import (
	"alta-test/entities"
	"alta-test/model"
	"alta-test/view"

	"github.com/jinzhu/copier"
	"github.com/labstack/gommon/log"
)

type ServiceModel struct {
	model model.ModelUser
}

func NewServiceModel(ModelUser model.ModelUser) *ServiceModel {
	return &ServiceModel{
		model: ModelUser,
	}
}

// CREATE NEW DATA USER
func (s *ServiceModel) CreateUser(newUser view.AddUser) (view.RespondUser, error) {
	var User entities.User
	copier.Copy(&User, &newUser)
	res, err := s.model.CreateUser(User)
	if err != nil {
		log.Warn(err)
		return view.RespondUser{}, err
	}
	var respond view.RespondUser
	copier.Copy(&respond, &res)
	return respond, nil
}

// GET ALL DATA USERS
func (s *ServiceModel) GetAllUsers() ([]view.RespondUser, error) {

	res, err := s.model.GetAllUsers()
	if err != nil {
		log.Warn(err)
		return []view.RespondUser{}, err
	}
	var respond []view.RespondUser
	copier.Copy(&respond, &res)
	return respond, nil
}

// GET DATA USER BY ID
func (s *ServiceModel) GetUserID(id uint) (view.RespondUser, error) {

	res, err := s.model.GetUserID(id)
	if err != nil {
		log.Warn(err)
		return view.RespondUser{}, err
	}
	var respond view.RespondUser
	copier.Copy(&respond, &res)
	return respond, nil
}

// UPDATE DATA USER BY ID
func (s *ServiceModel) UpdateUser(id uint, update view.UpdateUser) (view.RespondUser, error) {
	var User entities.User
	copier.Copy(&User, &update)
	res, err := s.model.UpdateUserID(id, User)
	if err != nil {
		log.Warn(err)
		return view.RespondUser{}, err
	}
	var respond view.RespondUser
	copier.Copy(&respond, &res)
	return respond, nil
}

// DELETE DATA USER BY ID
func (s *ServiceModel) DeleteUserID(id uint) error {

	err := s.model.DeleteUserID(id)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
