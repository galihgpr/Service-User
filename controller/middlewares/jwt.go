package middlewares

import (
	"alta-test/view"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
)

// Create Middleware Auth With JWT
func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check The Authorization, Is There Barier Auth or Empty Auth
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, view.StatusUnauthorized("Missing JWT token"))
			return
		}
		// Get Token JWT
		tokenString := authHeader[len("Bearer")+1:]

		// Check Token Is Valid Or Invalid.
		token, err := ParseToken(tokenString)
		fmt.Println(token)
		if err != nil {
			log.Warn(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, view.StatusUnauthorized(err.Error()))
			return
		}
		c.Next()
	}
}

// Generate New Token With Several Parameters
func GenerateToken(id int, name string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["role"] = role
	claims["expired"] = time.Now().Add(time.Hour * 48).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("SECRET_JWT"))
}

// Ectract Token To Get Data Id and Role Of User
func ExtractToken(c *gin.Context) (int, string) {
	authHeader := c.GetHeader("Authorization")
	token := authHeader[len("Bearer")+1:]
	tokenID, _ := ParseToken(token)
	if tokenID.Valid {
		claims := tokenID.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		role := claims["role"].(string)
		return id, role
	}
	return 0, ""
}

// Check Token Is Valid Or Not
func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])
		}
		return []byte("SECRET_JWT"), nil
	})

	if err != nil {
		log.Warn(err)
		return nil, errors.New("Invalid Token JWT")
	}

	return token, nil
}
