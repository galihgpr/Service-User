package handler

import "github.com/gin-gonic/gin"

type HandlerUser interface {
	CreateUser() gin.HandlerFunc
	GetAllUsers() gin.HandlerFunc
	GetUserID() gin.HandlerFunc
	UpdateUserID() gin.HandlerFunc
	DeleteUserID() gin.HandlerFunc
}
