package model

import (
	"alta-test/entities"
	"fmt"
	"strings"

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
	if err := d.db.Create(&newUser).Scan(&User).Error; err != nil {
		log.Warn(err)
		return entities.User{}, err
	}
	return User, nil
}

// GET ALL DATA USERS
func (d *ModelDB) GetAllUsers(name, role, sort string, sizePage, page int) ([]entities.User, error) {
	var AllUsers []entities.User

	filter := "SELECT * FROM users"
	if name != "" {
		filter += " WHERE LOWER(name) LIKE '%" + strings.ToLower(name) + "%'"
		if role != "" {
			filter += fmt.Sprintf(" AND role = '%s'", role)
		}
		if sort != "" {
			filter += " ORDER BY name " + sort
		}
		if sizePage != 0 {
			filter += fmt.Sprintf(" LIMIT %d", sizePage)
		}
		if page != 0 {
			filter += fmt.Sprintf(" OFFSET %d", page*sizePage)
		}
	} else if role != "" {
		filter += fmt.Sprintf(" WHERE role = '%s'", role)
		if sort != "" {
			filter += " ORDER BY name " + sort
		}
		if sizePage != 0 {
			filter += fmt.Sprintf(" LIMIT %d", sizePage)
		}
		if page != 0 {
			filter += fmt.Sprintf(" OFFSET %d", page*sizePage)
		}
	} else if sort != "" {
		filter += " ORDER BY name " + sort
		if sizePage != 0 {
			filter += fmt.Sprintf(" LIMIT %d", sizePage)
		}
		if page != 0 {
			filter += fmt.Sprintf(" OFFSET %d", page*sizePage)
		}
	} else if sizePage != 0 {
		filter += fmt.Sprintf(" LIMIT %d", sizePage)
		if page != 0 {
			filter += fmt.Sprintf(" OFFSET %d", page*sizePage)
		}
	}
	fmt.Println(filter)
	if err := d.db.Raw(filter).Find(&AllUsers).Error; err != nil {
		log.Warn(err)
		return []entities.User{}, err
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
