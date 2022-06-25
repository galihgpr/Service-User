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
			return
		}
		if err := u.valid.Struct(&newUser); err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("name, email, password, phone and rule are Required"))
			return
		}
		res, err := u.service.CreateUser(newUser)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("email have been regitered"))
			return
		}
		c.JSON(http.StatusCreated, view.StatusCreated(res, "Success create new data user"))
	}
}

// HANDLER GET ALL USER
func (u *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		role := c.Query("role")
		sort := c.Query("sort")
		size := c.Query("sizePage")
		pages := c.Query("page")
		if pages == "" {
			pages = "0"
		}
		if size == "" {
			size = "0"
		}
		fmt.Println(name, role, sort, "0", size, pages)
		sizePage, err := strconv.Atoi(size)
		if err != nil {
			fmt.Println(err)
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param sizePage must be integer"))
			return
		}

		page, err := strconv.Atoi(pages)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param page must be integer"))
			return
		}

		if page > 0 {
			page--
		}

		res, err := u.service.GetAllUsers(name, role, sort, sizePage, page)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound("Data User Not Found"))
			return
		}
		if res == nil {
			c.JSON(http.StatusNotFound, view.StatusNotFound("Data User Not Found"))
			return
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
			return
		}

		res, err := u.service.GetUserID(uint(userID))
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", userID)))
			return
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
			return
		}

		var updateUser view.UpdateUser

		if err := c.Bind(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, require JSON OR Form"))
			return
		}

		res, err := u.service.UpdateUserID(uint(userID), updateUser)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", userID)))
			return
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
			return
		}
		err = u.service.DeleteUserID(uint(userID))
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", userID)))
			return
		}
		c.JSON(http.StatusOK, view.StatusDeleted("Success delete data users"))
	}
}
