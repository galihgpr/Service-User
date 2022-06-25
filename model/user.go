package model

import (
	"alta-test/entities"
	"alta-test/view"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ModelDB struct {
	db *gorm.DB
}

func NewModelDB(DB *gorm.DB) *ModelDB {
	return &ModelDB{
		db: DB,
	}
}

// CREATE NEW DATA USER
func (d *ModelDB) CreateUser(newUser entities.User) (entities.User, error) {
	var User entities.User
	if err := d.db.Create(&newUser).Find(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// GET ALL DATA USERS
func (d *ModelDB) GetAllUsers() ([]view.RespondUser, error) {
	var AllUsers []view.RespondUser
	if err := d.db.Find(&AllUsers).Error; err != nil {
		log.Warn(err)
		return []view.RespondUser{}, err
	}
	return AllUsers, nil
}

// GET DATA USER BY ID
func (d *ModelDB) GetUserID(id uint) (entities.User, error) {
	var User entities.User
	if err := d.db.Where("id=?", id).First(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// UPADATE DATA USER BY ID AND
func (d *ModelDB) UpdateUserID(id uint, update entities.User) (entities.User, error) {
	var User entities.User
	if err := d.db.Where("id = ?", id).Updates(&update).Find(&User); err != nil {
		log.Warn(err)
	}
}
