package model

import (
	"alta-test/config"
	"alta-test/entities"
	"alta-test/view"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ModelDB struct {
	db *gorm.DB
}

func NewModelDB() *ModelDB {
	return &ModelDB{
		db: config.DB,
	}
}

// CREATE NEW DATA USER
func (d *ModelDB) CreateUser(newUser entities.User) (entities.User, error) {
	var User entities.User
	if err := d.db.Create(&newUser).Scan(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// GET ALL DATA USERS
func (d *ModelDB) GetAllUsers(name, role, sort string, sizePage, page int) ([]entities.User, error) {
	var AllUsers []entities.User

	// Page Must Be Decrement 1 For Valid Offset
	if page > 0 {
		page--
	}

	filter := "SELECT * FROM users"

	if name != "" || role != "" {
		filter += " WHERE "
	}

	// Check Filter
	if name != "" {
		filter += "LOWER(name) LIKE '%" + strings.ToLower(name) + "%' AND"
	}

	if role != "" {
		filter += fmt.Sprintf("role = '%s' ", role)
	}

	// Remove "AND" if role is empty
	filter = strings.TrimSuffix(filter, "AND")

	if sort != "" {
		filter += "ORDER BY name " + sort
	}

	if sizePage != 0 {
		filter += fmt.Sprintf("LIMIT %d ", sizePage)
	}

	if page != 0 {
		filter += fmt.Sprintf("OFFSET %d", page*sizePage)
	}

	if err := d.db.Raw(filter).Limit(sizePage).Offset(page * sizePage).Find(&AllUsers).Error; err != nil {
		log.Warn(err)
		return []entities.User{}, err
	}

	return AllUsers, nil
}

// GET DATA USER BY ID
func (d *ModelDB) GetUserID(id uint, role string) (entities.User, error) {
	filter := ""
	// User Only Access Data User With Role = "user"
	if role == "admin" {
		filter = fmt.Sprintf("id = %d", id)
	} else {
		filter = fmt.Sprintf("id = %d AND role = %s", id, role)
	}

	var User entities.User
	if err := d.db.Where(filter).First(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// UPADATE DATA USER BY ID AND
func (d *ModelDB) UpdateUserID(id uint, update entities.User) (entities.User, error) {
	var User entities.User
	if err := d.db.Where("id = ?", id).Updates(&update).First(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// DELETE DATA USER BY ID
func (d *ModelDB) DeleteUserID(id uint) error {
	var delete entities.User
	if err := d.db.Where("id = ?", id).First(&delete).Delete(&delete).Error; err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

// GET USER BY EMAIL && PASSWORD
func (d *ModelDB) GetUserLogin(login view.Login) (entities.User, error) {
	var User entities.User
	if err := d.db.Where("email = ? AND password = ?", login.Email, login.Password).First(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}
