package middlewares

import (
	"event-booking-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID int
	Email  string
}

func Authorization(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userInfo, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := User{
		UserID: userInfo.UserID,
		Email:  userInfo.Email,
	}

	c.Set("user", user)
	c.Next()
}
