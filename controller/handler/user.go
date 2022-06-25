package handler

import (
	"alta-test/service"
	"alta-test/view"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type UserHandler struct {
	service service.ServiceUser
	valid   *validator.Validate
}

func NewUserHandler(Service service.ServiceUser, Valid *validator.Validate) *UserHandler {
	return &UserHandler{
		service: Service,
		valid:   Valid,
	}
}

// HANDLER CREATE USER
func (u *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser view.AddUser
		if err := c.Bind(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, require JSON OR Form"))
		}
		if err := u.valid.Struct(&newUser); err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("name, email, password, phone and rule are Required"))
		}
		res, err := u.service.CreateUser(newUser)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusInternalServerError, view.StatusErrorServer())
		}
		c.JSON(http.StatusCreated, view.StatusCreated(res, "Success create new data user"))
	}
}

// HANDLER GET ALL USER
func (u *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := u.service.GetAllUsers()
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusInternalServerError, view.StatusErrorServer())
		}
		c.JSON(http.StatusOK, view.StatusSuccess(res, "Success get all data users"))
	}
}

// HANDLER GET USER BY ID
func (u *UserHandler) GetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
		}

		res, err := u.service.GetUserID(uint(userID))
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", id)))
		}

		c.JSON(http.StatusOK, view.StatusSuccess(res, "Success get data user"))
	}
}

// HANDLER UPDATE DATA USER ID
func (u *UserHandler) UpdateUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
		}

		var updateUser view.UpdateUser

		if err := c.Bind(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, require JSON OR Form"))
		}

		res, err := u.service.UpdateUserID(uint(userID), updateUser)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", id)))
		}

		c.JSON(http.StatusOK, view.StatusSuccess(res, "Success update data user"))
	}
}

// HANDLER DELETE USER ID
func (u *UserHandler) DeleteUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
		}
		err = u.service.DeleteUserID(uint(userID))
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", id)))
		}
		c.JSON(http.StatusOK, view.StatusDeleted("Success delete data users"))
	}
}
