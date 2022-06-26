package service

import (
	"alta-test/entities"
	mocks "alta-test/mocks/model"
	"alta-test/view"
	"errors"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var MockUser = []entities.User{
	{
		Model:    gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     "Galih",
		Email:    "galih@gmail.com",
		Password: "galih",
		Phone:    "123456",
		Role:     "admin",
	},
	{
		Model:    gorm.Model{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Name:     "Putri",
		Email:    "putri@gmail.com",
		Password: "putri",
		Phone:    "654321",
		Role:     "user",
	},
}

// Test Service Create User
func TestCreateUser(t *testing.T) {
	var NewUser view.AddUser
	copier.Copy(&NewUser, &MockUser[0])
	userRepo := mocks.NewModelUser(t)
	t.Run("Success Create User", func(t *testing.T) {
		userRepo.On("CreateUser", mock.Anything).Return(MockUser[0], nil).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.CreateUser(NewUser)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res.Name)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Phone, res.Phone)
		assert.Equal(t, MockUser[0].Role, res.Role)
	})
	t.Run("Error Create User", func(t *testing.T) {
		userRepo.On("CreateUser", mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.CreateUser(NewUser)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.Name)
		assert.NotEqual(t, MockUser[0].Email, res.Email)
		assert.NotEqual(t, MockUser[0].Phone, res.Phone)
		assert.NotEqual(t, MockUser[0].Role, res.Role)
	})
}

// Test Service Get All Users
func TestGetAllUsers(t *testing.T) {

	userRepo := mocks.NewModelUser(t)
	name := "g"
	role := ""
	sort := ""
	pageSize := 0
	page := 0

	t.Run("Success Get All User", func(t *testing.T) {
		userRepo.On("GetAllUsers", name, role, sort, pageSize, page).Return(MockUser, nil).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetAllUsers(name, role, sort, pageSize, page)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res[0].Name)
		assert.Equal(t, MockUser[0].Email, res[0].Email)
		assert.Equal(t, MockUser[0].Phone, res[0].Phone)
		assert.Equal(t, MockUser[0].Role, res[0].Role)
	})
	t.Run("Error Get All User", func(t *testing.T) {
		userRepo.On("GetAllUsers", name, role, sort, pageSize, page).Return(nil, errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetAllUsers(name, role, sort, pageSize, page)
		assert.Error(t, err)
		assert.Equal(t, []view.RespondUser{}, res)
	})

}

// Test Service Get User ID
func TestGetUserID(t *testing.T) {
	var User view.RespondUser
	copier.Copy(&User, &MockUser[0])
	userRepo := mocks.NewModelUser(t)

	t.Run("Success Get User ID", func(t *testing.T) {
		userRepo.On("GetUserID", mock.Anything, "admin").Return(MockUser[0], nil).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetUserID(uint(1), "admin")
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res.Name)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Phone, res.Phone)
		assert.Equal(t, MockUser[0].Role, res.Role)
	})
	t.Run("Error Get User ID", func(t *testing.T) {
		userRepo.On("GetUserID", mock.Anything, "admin").Return(entities.User{}, errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetUserID(uint(1), "admin")
		assert.Error(t, err)
		assert.Equal(t, view.RespondUser{}, res)
	})

}

// Test Update User ID
func TestUpdateUserID(t *testing.T) {
	var Update view.UpdateUser
	copier.Copy(&Update, &MockUser[0])
	userRepo := mocks.NewModelUser(t)
	t.Run("Success Update User", func(t *testing.T) {
		userRepo.On("UpdateUserID", uint(1), mock.Anything).Return(MockUser[0], nil).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.UpdateUserID(uint(1), Update)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res.Name)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Phone, res.Phone)
		assert.Equal(t, MockUser[0].Role, res.Role)
	})
	t.Run("Error Update User", func(t *testing.T) {
		userRepo.On("UpdateUserID", uint(1), mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.UpdateUserID(uint(1), Update)
		assert.Error(t, err)
		assert.NotEqual(t, MockUser[0].Name, res.Name)
		assert.NotEqual(t, MockUser[0].Email, res.Email)
		assert.NotEqual(t, MockUser[0].Phone, res.Phone)
		assert.NotEqual(t, MockUser[0].Role, res.Role)
	})
}

// Test Service Delete User ID
func TestDeleteUserID(t *testing.T) {
	var User view.RespondUser
	copier.Copy(&User, &MockUser[0])
	userRepo := mocks.NewModelUser(t)

	t.Run("Success Delete User ID", func(t *testing.T) {
		userRepo.On("DeleteUserID", mock.Anything).Return(nil).Once()
		userService := NewServiceModel(userRepo)
		err := userService.DeleteUserID(uint(1))
		assert.NoError(t, err)
	})
	t.Run("Error Delete User ID", func(t *testing.T) {
		userRepo.On("DeleteUserID", mock.Anything).Return(errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		err := userService.DeleteUserID(uint(1))
		assert.Error(t, err)
	})

}

// Test Service Delete
func TestGetUserLogin(t *testing.T) {
	var User view.Login
	copier.Copy(&User, &MockUser[0])
	userRepo := mocks.NewModelUser(t)

	t.Run("Success Login", func(t *testing.T) {
		userRepo.On("GetUserLogin", mock.Anything).Return(MockUser[0], nil).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetUserLogin(User)
		assert.NoError(t, err)
		assert.Equal(t, MockUser[0].Name, res.Name)
		assert.Equal(t, MockUser[0].Email, res.Email)
		assert.Equal(t, MockUser[0].Phone, res.Phone)
		assert.Equal(t, MockUser[0].Role, res.Role)
	})
	t.Run("Error Get User Login", func(t *testing.T) {
		userRepo.On("GetUserLogin", mock.Anything).Return(entities.User{}, errors.New("Error Access Database")).Once()
		userService := NewServiceModel(userRepo)
		res, err := userService.GetUserLogin(User)
		assert.Error(t, err)
		assert.Equal(t, view.RespondUser{}, res)
	})

}
