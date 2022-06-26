package handler

import (
	"alta-test/controller/middlewares"
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

// HANDLER LOGIN USER
func (u *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Binding Login Data
		var login view.Login
		if err := c.Bind(&login); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, JSON OR Form is required"))
			return
		}

		// Validate Data
		if err := u.valid.Struct(&login); err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Email and password are Required"))
			return
		}

		// Get Data User From Service
		res, err := u.service.GetUserLogin(login)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Email or password is wrong"))
			return
		}

		// Generate Token With JWT
		token, err := middlewares.GenerateToken(int(res.ID), res.Name, res.Role)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Error create token"))
			return
		}

		// Create Respond Login
		response := view.RespondLogin{
			Token: token,
			User:  res,
		}
		c.JSON(http.StatusOK, view.StatusSuccess(response, "Success login"))
	}
}

// HANDLER CREATE USER
func (u *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract Token to Check Role of User
		_, roleUser := middlewares.ExtractToken(c)
		if roleUser != "admin" {
			c.JSON(http.StatusUnauthorized, view.StatusUnauthorized("Access denied, only admin can access"))
			return
		}

		// Binding Create Data User
		var newUser view.AddUser
		if err := c.Bind(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, JSON OR Form is required"))
			return
		}

		// Validate Data
		if err := u.valid.Struct(&newUser); err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Name, email, password, phone and rule are required"))
			return
		}

		// Access Service to Create User
		res, err := u.service.CreateUser(newUser)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Email have been registered"))
			return
		}
		c.JSON(http.StatusCreated, view.StatusCreated(res, "Success create new data user"))
	}
}

// HANDLER GET ALL USER
func (u *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Query Param For Filter Data User
		name := c.Query("name")
		role := c.Query("role")
		sort := c.Query("sort")
		size := c.Query("sizePage")
		pages := c.Query("page")

		// Change Value Query Pages and SixePagez to 0 If Query Value is Blank Data
		// It For In The Next Step, Convert String To Int Didn't Error
		if pages == "" {
			pages = "0"
		}
		if size == "" {
			size = "0"
		}

		// Extract Token to Check Role of User
		_, roleUser := middlewares.ExtractToken(c)
		if roleUser != "admin" && roleUser != "user" {
			c.JSON(http.StatusUnauthorized, view.StatusUnauthorized("Access denied, only admin & user can access"))
			return
		}

		// Filter Role Active With Value "user", If userRole = "user"
		if roleUser == "user" {
			role = "user"

		}
		// Convert SizePage And Page From String To Int
		sizePage, err := strconv.Atoi(size)
		if err != nil {
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

		// Access Service to Get All Data User
		res, err := u.service.GetAllUsers(name, role, sort, sizePage, page)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound("Data user not found"))
			return
		}
		if res == nil {
			c.JSON(http.StatusNotFound, view.StatusNotFound("Data user not found"))
			return
		}
		c.JSON(http.StatusOK, view.StatusSuccess(res, "Success get all data users"))
	}
}

// HANDLER GET USER BY ID
func (u *UserHandler) GetUserID() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract Token to Check Role of User
		_, roleUser := middlewares.ExtractToken(c)
		if roleUser != "admin" && roleUser != "user" {
			c.JSON(http.StatusUnauthorized, view.StatusUnauthorized("Access denied, only admin & user can access"))
			return
		}

		// Get Param Value Of ID
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
			return
		}

		// Access Service to Get User ID
		res, err := u.service.GetUserID(uint(userID), roleUser)
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

		// Extract Token to Check Role of User
		_, roleUser := middlewares.ExtractToken(c)
		if roleUser != "admin" {
			c.JSON(http.StatusUnauthorized, view.StatusUnauthorized("Access denied, only admin can access"))
			return
		}

		// Get Param Value Of ID
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
			return
		}

		// Binding Update Data User
		var updateUser view.UpdateUser
		if err := c.Bind(&updateUser); err != nil {
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Wrong data input format, JSON OR Form is required"))
			return
		}

		// Access Service To Update Data User
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

		// Extract Token to Check Role of User
		_, roleUser := middlewares.ExtractToken(c)
		if roleUser != "admin" {
			c.JSON(http.StatusUnauthorized, view.StatusUnauthorized("Access denied, only admin can access"))
			return
		}

		// Get Param Value Of ID
		id := c.Param("id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusBadRequest, view.StatusBadRequest("Param id must be integer"))
			return
		}

		// Access Service To Delete Data User
		err = u.service.DeleteUserID(uint(userID))
		if err != nil {
			log.Warn(err)
			c.JSON(http.StatusNotFound, view.StatusNotFound(fmt.Sprintf("User with ID %d not found", userID)))
			return
		}
		c.JSON(http.StatusOK, view.StatusDeleted("Success delete data users"))
	}
}
