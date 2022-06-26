package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespondUser struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
	Role  string `json:"role" form:"role"`
}

type RespondLogin struct {
	Token string `json:"token"`
	User  RespondUser
}

// HANDLING RESPONSE STATUS SUCCESS
func StatusCreated(data interface{}, message string) gin.H {
	return gin.H{
		"code":    http.StatusCreated,
		"message": message,
		"data":    data,
	}
}

func StatusSuccess(data interface{}, message string) gin.H {
	return gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	}
}

func StatusDeleted(message string) gin.H {
	return gin.H{
		"code":    http.StatusOK,
		"message": message,
	}
}

// HANDLING RESPONSE STATUS ERROR
func StatusBadRequest(message string) gin.H {
	return gin.H{
		"code":    http.StatusBadRequest,
		"message": message,
	}
}

func StatusNotFound(message string) gin.H {
	return gin.H{
		"code":    http.StatusNotFound,
		"message": message,
	}
}

func StatusUnauthorized(message string) gin.H {
	return gin.H{
		"code":    http.StatusUnauthorized,
		"message": message,
	}
}
